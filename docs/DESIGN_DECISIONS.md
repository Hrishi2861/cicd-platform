# Design Decisions & Lessons Learned

This document captures the key architectural choices made while building this
platform, the alternatives that were considered, and the tradeoffs behind each
decision. The goal is to make the reasoning explicit rather than presenting the
final state as if it were inevitable.

---

## Why EKS over self-managed Kubernetes

**Decision:** Use Amazon EKS (managed control plane) instead of running the
Kubernetes control plane on self-managed EC2 (e.g. kubeadm or kops).

**Reasoning:**

- **Control-plane operations are undifferentiated heavy lifting.** etcd backups,
  API server HA, certificate rotation, and version upgrades are exactly the kind
  of work that adds no product value but carries a lot of operational risk. EKS
  manages the control plane across multiple AZs for a flat hourly fee.
- **IAM integration.** EKS integrates natively with AWS IAM (IRSA — IAM Roles for
  Service Accounts), which lets workloads assume least-privilege roles without
  long-lived credentials. Wiring this up on a self-managed cluster is possible but
  fiddly.
- **Upgrade path.** Managed node groups + control-plane upgrades reduce a
  multi-day, error-prone task to a mostly automated one.

**Tradeoffs accepted:**

- Less control over control-plane flags and API server configuration.
- Vendor lock-in to AWS (mitigated by keeping all workload manifests portable —
  plain Helm/Kustomize, no AWS-specific CRDs in the app layer).
- Cost: the control plane has a fixed hourly cost even when idle. For a portfolio
  or small workload this is the main downside.

**When I'd revisit:** a multi-cloud requirement, or a need for custom control-plane
tuning, would push me back toward self-managed or a different distribution (k3s,
Talos).

---

## Why ArgoCD over Flux

**Decision:** Use ArgoCD for GitOps continuous delivery.

**Reasoning:**

- **The UI is a real feature, not a nicety.** Being able to visualize the
  application topology, see live vs. desired state diffs, and trigger/inspect syncs
  from a dashboard is genuinely useful when demonstrating GitOps and when
  debugging drift. Flux is excellent but is CLI/CRD-first.
- **App-of-apps and ApplicationSets** map cleanly onto a multi-service repo like
  this one.
- **Self-heal + prune** give the declarative, drift-correcting behavior I wanted
  with a single `syncPolicy` block.

**Tradeoffs accepted:**

- ArgoCD is a heavier install than Flux and introduces its own RBAC/SSO surface to
  manage.
- Flux's native integration with the Kustomize/Helm controllers is arguably more
  "Kubernetes-idiomatic" and composes better with a pure controller model.

**When I'd revisit:** if I wanted a lighter footprint, tighter Terraform-driven
bootstrapping, or a pure-controller model with no extra UI to secure, Flux would be
the stronger choice.

---

## Why Kustomize overlays instead of separate Helm values files

**Decision:** Use Helm charts for the *base* packaging of each service, and
Kustomize overlays for *environment* differences (dev / staging / prod) rather
than maintaining `values-dev.yaml`, `values-staging.yaml`, `values-prod.yaml`.

**Reasoning:**

- **Overlays express intent as patches, not full copies.** A prod overlay says
  "replicas: 3, add resource limits" as a small strategic-merge/JSON patch. A
  separate values file tends to drift into a near-complete duplicate of the base,
  and it's hard to see *what actually differs* between environments at a glance.
- **`kubectl apply -k` needs no extra tooling** in-cluster, and the overlays are
  readable as plain diffs in review.
- **Separation of concerns:** Helm answers "how is this service templated and
  packaged?"; Kustomize answers "how does this environment differ?". Keeping those
  two axes in different tools kept each one simpler.

**Tradeoffs accepted:**

- Two templating systems in one repo is more concepts for a newcomer to learn.
- Rendering Helm output and then patching with Kustomize can be awkward; the split
  works cleanly here because the environment deltas (replicas, image tag,
  resources) are small and structural.

**When I'd revisit:** if environment differences grew to include large,
value-driven configuration (feature flags, many tunables), Helm values files —
or a values-per-env approach with a single tool — would reduce the cognitive
overhead.

---

## Tradeoffs: RDS vs. an in-cluster database

**Decision:** Use Amazon RDS (PostgreSQL) for persistent data rather than running
PostgreSQL inside the cluster (e.g. a StatefulSet or an operator like CNPG/Zalando).

**Reasoning:**

- **Stateful workloads are the hardest thing to run well on Kubernetes.** Backups,
  point-in-time recovery, failover, storage resizing, and minor-version patching
  are all solved problems on RDS and all sharp edges in-cluster.
- **Blast radius.** Keeping the database off the cluster means a bad rollout, a
  node failure, or a cluster upgrade can't take data down with it.
- **Security.** RDS lives in private subnets, reachable only through a dedicated
  security group, with encryption at rest — a clean network isolation story that
  Terraform expresses declaratively.

**Tradeoffs accepted:**

- Higher cost than a single in-cluster Postgres pod, and another AWS-managed
  service to provision.
- Local development still uses an in-cluster/containerized Postgres (via
  `docker-compose`) for speed and zero cost — so there's a deliberate dev/prod
  parity gap at the database layer. The schema (`scripts/init-db.sql`) is shared to
  keep that gap small.

**When I'd revisit:** a strict "no managed services" or cost-minimization
constraint, or a need for a database flavor RDS doesn't offer, would make a
well-run in-cluster operator (e.g. CloudNativePG) worth the operational cost.

---

## Smaller decisions worth noting

- **Go for the services.** Small static binaries, fast cold starts, and
  multi-stage Docker builds that produce tiny final images — a good fit for
  container-first microservices.
- **API Gateway pattern.** A single ingress point (`api-gateway`) that fans out to
  `user-service` and `order-service` keeps routing/observability concerns in one
  place and lets the backend services stay focused.
- **Prometheus + Grafana over a hosted APM.** Keeps the stack self-contained,
  portable, and free to run locally; every service exposes `/metrics` and the
  middleware records request metrics uniformly.
- **Trivy in CI.** Filesystem and dependency scanning runs on every push so
  vulnerabilities surface at PR time, not in production.

---

## Lessons learned

- **Commit history is documentation.** Building this in a logical order — services,
  then containers, then infra, then delivery, then observability — made each layer
  easier to reason about and would make it far easier for a collaborator to follow.
- **Keep the two templating tools doing one job each.** The moment Helm values and
  Kustomize patches both tried to own environment config, things got confusing.
  Drawing the line at "Helm packages, Kustomize environments" removed that friction.
- **Dev/prod parity is a spectrum, not a binary.** Using containerized Postgres
  locally and RDS in prod is a pragmatic tradeoff; sharing the schema and keeping
  connection handling identical kept the gap manageable.
- **Managed services are worth it for state.** The time not spent operating etcd
  and Postgres went into the parts of the platform that actually demonstrate
  engineering: the CI/CD pipeline, GitOps flow, and observability.
