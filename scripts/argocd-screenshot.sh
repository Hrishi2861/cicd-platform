#!/usr/bin/env bash
set -euo pipefail

# ──────────────────────────────────────────────────────────────────────────────
#  argocd-screenshot.sh — Spin up a local kind cluster + ArgoCD + api-gateway
#  then print instructions to screenshot the Synced / Healthy state.
#
#  Usage:
#    ./scripts/argocd-screenshot.sh              # full interactive flow
#    ./scripts/argocd-screenshot.sh --auto        # auto-approve installs
#    ./scripts/argocd-screenshot.sh --destroy     # tear down the kind cluster
# ──────────────────────────────────────────────────────────────────────────────

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
BINDIR="${HOME}/.local/bin"
KUBECTL="${BINDIR}/kubectl"
KIND="${BINDIR}/kind"
CLUSTER_NAME="cicd-platform"
NAMESPACE="argocd"
TARGET_NS="cicd-platform"
APP_YAML="${ROOT}/infra/argocd/applications/api-gateway.yaml"

# ── colours ───────────────────────────────────────────────────────────────────
RED='\033[0;31m'; GREEN='\033[0;32m'; YELLOW='\033[1;33m'
CYAN='\033[0;36m'; BOLD='\033[1m'; NC='\033[0m'
info()  { echo -e "${CYAN}${BOLD}[INFO]${NC}  $*"; }
ok()    { echo -e "${GREEN}${BOLD}[OK]${NC}    $*"; }
warn()  { echo -e "${YELLOW}${BOLD}[WARN]${NC}  $*"; }
err()   { echo -e "${RED}${BOLD}[ERROR]${NC} $*"; }

# ── helpers ───────────────────────────────────────────────────────────────────
prompt_confirm() {
  local msg="$1"
  local answer
  echo -en "${YELLOW}${msg} [Y/n] ${NC}"
  read -r answer
  [[ -z "$answer" || "$answer" =~ ^[Yy] ]]
}

command_exists() { command -v "$1" &>/dev/null; }

ensure_bin_dir() {
  mkdir -p "$BINDIR"
}

# ── step: install kind ────────────────────────────────────────────────────────
install_kind() {
  ensure_bin_dir
  if [ -x "$KIND" ]; then
    ok "kind already installed at $KIND ($("$KIND" version))"
    return
  fi
  info "Downloading kind ..."
  curl -fsSLo "$KIND" "https://kind.sigs.k8s.io/dl/v0.27.0/kind-linux-amd64"
  chmod +x "$KIND"
  ok "kind installed at $KIND"
}

# ── step: install kubectl ─────────────────────────────────────────────────────
install_kubectl() {
  ensure_bin_dir
  if [ -x "$KUBECTL" ]; then
    ok "kubectl already installed at $KUBECTL ($("$KUBECTL" version --client --short 2>/dev/null || true))"
    return
  fi
  info "Downloading kubectl ..."
  curl -fsSLo "$KUBECTL" "https://dl.k8s.io/release/v1.32.0/bin/linux/amd64/kubectl"
  chmod +x "$KUBECTL"
  ok "kubectl installed at $KUBECTL"
}

# ── step: create kind cluster ─────────────────────────────────────────────────
create_cluster() {
  if "$KIND" get clusters 2>/dev/null | grep -q "^${CLUSTER_NAME}$"; then
    ok "kind cluster '${CLUSTER_NAME}' already exists"
    return
  fi
  info "Creating kind cluster '${CLUSTER_NAME}' (this pulls a ~1 GB node image) ..."
  # Use K8s 1.31 — newer versions have stricter CRD annotation limits that conflict with ArgoCD's install manifests
  "$KIND" create cluster --name "$CLUSTER_NAME" --image kindest/node:v1.31.0
  ok "kind cluster created"
}

# ── step: load image ──────────────────────────────────────────────────────────
load_image() {
  # docker-compose builds the image as cicd-platform-api-gateway:latest,
  # but the Helm chart references api-gateway:latest.  Tag it so kind
  # has the exact name the chart expects.
  local src="cicd-platform-api-gateway:latest"
  local dst="api-gateway:latest"

  if ! docker image inspect "$src" &>/dev/null; then
    err "Image $src not found locally. Run 'make dev-up' or 'docker compose build' first."
    return 1
  fi

  docker tag "$src" "$dst" 2>/dev/null || true

  if "$KIND" load docker-image "$dst" --name "$CLUSTER_NAME" 2>&1 | grep -q "node(s) already have image"; then
    ok "image $dst already loaded into cluster"
    return
  fi

  info "Loading $dst (from $src) into kind cluster ..."
  "$KIND" load docker-image "$dst" --name "$CLUSTER_NAME"
  ok "$dst loaded"
}

# ── step: install ArgoCD ─────────────────────────────────────────────────────
install_argocd() {
  if "$KUBECTL" -n "$NAMESPACE" get deployment argocd-server &>/dev/null 2>&1; then
    ok "ArgoCD already installed"
    return
  fi
  info "Installing ArgoCD (this pulls ~few hundred MB) ..."
  "$KUBECTL" create namespace "$NAMESPACE" --dry-run=client -o yaml | "$KUBECTL" apply -f -
  # v2.12.6 is the last version without oversized CRD annotations that fail on newer K8s
  "$KUBECTL" apply -n "$NAMESPACE" -f "https://raw.githubusercontent.com/argoproj/argo-cd/v2.12.6/manifests/install.yaml"
  # apply the repo's custom ConfigMap on top
  "$KUBECTL" apply -f "${ROOT}/infra/argocd/install.yaml"
  ok "ArgoCD installed"
}

# ── step: wait for ArgoCD ─────────────────────────────────────────────────────
wait_argocd() {
  info "Waiting for ArgoCD components to be ready (this can take a few minutes) ..."
  "$KUBECTL" -n "$NAMESPACE" wait --for=condition=Available deployment/argocd-server --timeout=180s
  "$KUBECTL" -n "$NAMESPACE" wait --for=condition=Available deployment/argocd-redis --timeout=120s 2>/dev/null || true
  "$KUBECTL" -n "$NAMESPACE" wait --for=condition=Available deployment/argocd-repo-server --timeout=120s 2>/dev/null || true
  ok "ArgoCD is ready"
}

# ── step: apply api-gateway Application ───────────────────────────────────────
apply_application() {
  if "$KUBECTL" -n "$NAMESPACE" get application api-gateway &>/dev/null 2>&1; then
    ok "api-gateway Application already exists, reapplying ..."
  fi
  "$KUBECTL" create namespace "$TARGET_NS" --dry-run=client -o yaml | "$KUBECTL" apply -f -
  "$KUBECTL" apply -f "$APP_YAML"
  ok "api-gateway Application applied"
}

# ── step: wait for sync ───────────────────────────────────────────────────────
wait_sync() {
  info "Waiting for api-gateway to sync and become healthy ..."
  local timeout=120
  local waited=0
  while [ $waited -lt $timeout ]; do
    local status
    status="$("$KUBECTL" -n "$NAMESPACE" get application api-gateway -o jsonpath='{.status.health.status}' 2>/dev/null || echo '')"
    if [ "$status" = "Healthy" ]; then
      ok "api-gateway is Synced and Healthy!"
      return 0
    fi
    sleep 5
    waited=$((waited + 5))
    echo -n "."
  done
  echo ""
  warn "Timed out waiting for Healthy status. Check: kubectl -n argocd get application api-gateway -o yaml"
  return 1
}

# ── step: show instructions ───────────────────────────────────────────────────
show_instructions() {
  local pass
  pass="$("$KUBECTL" -n "$NAMESPACE" get secret argocd-initial-admin-secret -o jsonpath='{.data.password}' 2>/dev/null | base64 -d || echo '<secret-not-yet-available>')"
  echo ""
  echo -e "${GREEN}${BOLD}══════════════════════════════════════════════════════════════${NC}"
  echo -e "${GREEN}${BOLD}  ArgoCD is ready — time to take the screenshot!${NC}"
  echo -e "${GREEN}${BOLD}══════════════════════════════════════════════════════════════${NC}"
  echo ""
  echo -e "  ${BOLD}1. Port-forward the ArgoCD UI:${NC}"
  echo "     kubectl port-forward svc/argocd-server -n argocd 8080:443"
  echo ""
  echo -e "  ${BOLD}2. Open in your browser:${NC}"
  echo "     https://localhost:8080"
  echo ""
  echo -e "  ${BOLD}3. Login:${NC}"
  echo "     Username: admin"
  echo "     Password: ${pass}"
  echo ""
  echo -e "  ${BOLD}4. Look at the api-gateway Application:${NC}"
  echo "     It should show green ${GREEN}Synced${NC} + ${GREEN}Healthy${NC} badges."
  echo ""
  echo -e "  ${BOLD}5. Take a screenshot:${NC}"
  echo "     Save it as ${ROOT}/docs/screenshots/argocd-sync.png"
  echo ""
  echo -e "  ${BOLD}6. When done, tear down:${NC}"
  echo "     $0 --destroy"
  echo ""
  echo -e "${GREEN}${BOLD}══════════════════════════════════════════════════════════════${NC}"
}

# ── destroy ───────────────────────────────────────────────────────────────────
destroy() {
  info "Destroying kind cluster '${CLUSTER_NAME}' ..."
  "$KIND" delete cluster --name "$CLUSTER_NAME" 2>/dev/null || true
  ok "Cluster destroyed"
  exit 0
}

# ── main ──────────────────────────────────────────────────────────────────────
main() {
  echo -e "${CYAN}${BOLD}"
  echo "  ╔══════════════════════════════════════════════════════════╗"
  echo "  ║       ArgoCD Screenshot Setup                            ║"
  echo "  ║       ${ROOT}  ║"
  echo "  ╚══════════════════════════════════════════════════════════╝"
  echo -e "${NC}"

  # ── preflight: docker ───────────────────────────────────────────────────────
  if ! command_exists docker; then
    err "Docker is required but not found. Install docker first."
    exit 1
  fi
  if ! docker info &>/dev/null; then
    err "Docker daemon is not running."
    exit 1
  fi

  # ── ensure PATH includes BINDIR ─────────────────────────────────────────────
  export PATH="${BINDIR}:${PATH}"

  # ── run steps ───────────────────────────────────────────────────────────────
  AUTO=false
  for arg in "$@"; do
    case "$arg" in
      --destroy) destroy ;;
      --auto) AUTO=true ;;
    esac
  done

  install_kind
  install_kubectl
  create_cluster
  load_image
  install_argocd
  wait_argocd
  apply_application
  wait_sync
  show_instructions
}

main "$@"
