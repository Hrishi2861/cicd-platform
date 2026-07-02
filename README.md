<p align="center">
  <h1 align="center">Cloud-Native CI/CD Platform</h1>
  <p align="center">Production-ready microservices platform with full CI/CD, GitOps & Observability</p>
</p>

<p align="center">
  <a href="https://golang.org/">
    <img src="https://img.shields.io/badge/Go-1.21-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go"/>
  </a>
  <a href="https://www.docker.com/">
    <img src="https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white" alt="Docker"/>
  </a>
  <a href="https://kubernetes.io/">
    <img src="https://img.shields.io/badge/Kubernetes-326CE5?style=for-the-badge&logo=kubernetes&logoColor=white" alt="Kubernetes"/>
  </a>
  <a href="https://www.terraform.io/">
    <img src="https://img.shields.io/badge/Terraform-7B42BC?style=for-the-badge&logo=terraform&logoColor=white" alt="Terraform"/>
  </a>
  <a href="https://aws.amazon.com/">
    <img src="https://img.shields.io/badge/AWS-232F3E?style=for-the-badge&logo=amazon-aws&logoColor=white" alt="AWS"/>
  </a>
</p>

<p align="center">
  <a href="https://github.com/features/actions">
    <img src="https://img.shields.io/badge/GitHub_Actions-2088FF?style=for-the-badge&logo=github-actions&logoColor=white" alt="GitHub Actions"/>
  </a>
  <a href="https://argoproj.github.io/cd/">
    <img src="https://img.shields.io/badge/ArgoCD-1e40af?style=for-the-badge&logo=argo&logoColor=white" alt="ArgoCD"/>
  </a>
  <a href="https://helm.sh/">
    <img src="https://img.shields.io/badge/Helm-0F1689?style=for-the-badge&logo=helm&logoColor=white" alt="Helm"/>
  </a>
  <a href="https://kustomize.io/">
    <img src="https://img.shields.io/badge/Kustomize-6B55C6?style=for-the-badge&logo=kustomize&logoColor=white" alt="Kustomize"/>
  </a>
  <a href="https://www.postgresql.org/">
    <img src="https://img.shields.io/badge/PostgreSQL-4169E1?style=for-the-badge&logo=postgresql&logoColor=white" alt="PostgreSQL"/>
  </a>
</p>

<p align="center">
  <a href="https://prometheus.io/">
    <img src="https://img.shields.io/badge/Prometheus-E6522C?style=for-the-badge&logo=prometheus&logoColor=white" alt="Prometheus"/>
  </a>
  <a href="https://grafana.com/">
    <img src="https://img.shields.io/badge/Grafana-F46800?style=for-the-badge&logo=grafana&logoColor=white" alt="Grafana"/>
  </a>
  <a href="https://aquasecurity.github.io/trivy/">
    <img src="https://img.shields.io/badge/Trivy-2F77C8?style=for-the-badge&logo=aqua&logoColor=white" alt="Trivy"/>
  </a>
  <a href="https://github.com/aquasecurity/trivy">
    <img src="https://img.shields.io/badge/Security_Scan-E04F5F?style=for-the-badge&logo=dependabot&logoColor=white" alt="Security"/>
  </a>
  <a href="https://opensource.org/licenses/MIT">
    <img src="https://img.shields.io/badge/License-MIT-green?style=for-the-badge" alt="License"/>
  </a>
</p>

---

## Overview

A complete **DevOps & Cloud Engineering** platform demonstrating enterprise-grade practices for building, deploying, and operating cloud-native applications. From code commit to production deployment — fully automated.

### Key Features

- **Microservices Architecture** — API Gateway pattern with 3 independently deployable services
- **End-to-End CI/CD** — Automated testing, security scanning, and container builds
- **GitOps** — Declarative, self-healing deployments with ArgoCD
- **Infrastructure as Code** — Complete AWS provisioning with Terraform modules
- **Observability** — Real-time metrics, dashboards, and intelligent alerting
- **Multi-Environment** — Dev / Staging / Prod with Kustomize overlays

---

## Architecture

![Architecture](docs/architecture.svg)

---

## Demo

<!-- Drop the recording in once captured. -->
<!-- ![Demo](docs/screenshots/demo.gif) -->

> A short end-to-end demo — `argocd app sync` promoting a change, followed by a
> walkthrough of the live Grafana dashboard — is the best way to show this
> platform in motion.

**How to record one:**

```bash
# Option A — asciinema (terminal only, lightweight)
asciinema rec docs/screenshots/demo.cast
#  ... run:  kubectl apply -f infra/argocd/applications.yaml
#            argocd app sync api-gateway user-service order-service
#            argocd app wait api-gateway --health
#  Ctrl-D to stop, then convert to a GIF:
agg docs/screenshots/demo.cast docs/screenshots/demo.gif

# Option B — full-screen capture (to include the Grafana + ArgoCD UIs)
#  Use a screen recorder, walk through:
#    1. argocd app sync in the terminal
#    2. the ArgoCD UI showing Synced / Healthy
#    3. the Grafana Microservices dashboard with live traffic
#  Export to docs/screenshots/demo.gif and uncomment the image tag above.
```

## Screenshots

Proof artifacts live in [`docs/screenshots/`](docs/screenshots/) — see that
folder's README for exactly what to capture. Once added they render here:

### Grafana Dashboard

![Grafana](docs/screenshots/grafana-dashboard.png)

### ArgoCD Sync

![ArgoCD](docs/screenshots/argocd-sync.png)

### GitHub Actions

![Actions](docs/screenshots/github-actions-run.png)

---

## Tech Stack

### Core

| Badge | Technology | Purpose |
|-------|------------|---------|
| <img src="https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=go&logoColor=white" alt="Go"/> | **Go 1.21** | Microservices runtime |
| <img src="https://img.shields.io/badge/Docker-2496ED?style=flat-square&logo=docker&logoColor=white" alt="Docker"/> | **Docker** | Containerization with multi-stage builds |
| <img src="https://img.shields.io/badge/Kubernetes-326CE5?style=flat-square&logo=kubernetes&logoColor=white" alt="K8s"/> | **Kubernetes** | Container orchestration |

### Infrastructure

| Badge | Technology | Purpose |
|-------|------------|---------|
| <img src="https://img.shields.io/badge/Terraform-7B42BC?style=flat-square&logo=terraform&logoColor=white" alt="Terraform"/> | **Terraform** | Infrastructure as Code |
| <img src="https://img.shields.io/badge/AWS-232F3E?style=flat-square&logo=amazon-aws&logoColor=white" alt="AWS"/> | **AWS** | VPC, EKS, RDS, ECR |
| <img src="https://img.shields.io/badge/Helm-0F1689?style=flat-square&logo=helm&logoColor=white" alt="Helm"/> | **Helm** | Kubernetes package manager |
| <img src="https://img.shields.io/badge/Kustomize-6B55C6?style=flat-square&logo=kustomize&logoColor=white" alt="Kustomize"/> | **Kustomize** | Environment-specific configs |

### CI/CD & GitOps

| Badge | Technology | Purpose |
|-------|------------|---------|
| <img src="https://img.shields.io/badge/GitHub_Actions-2088FF?style=flat-square&logo=github-actions&logoColor=white" alt="Actions"/> | **GitHub Actions** | CI/CD pipeline |
| <img src="https://img.shields.io/badge/ArgoCD-1e40af?style=flat-square&logo=argo&logoColor=white" alt="ArgoCD"/> | **ArgoCD** | GitOps continuous delivery |
| <img src="https://img.shields.io/badge/GHCR-2088FF?style=flat-square&logo=github&logoColor=white" alt="GHCR"/> | **GitHub CR** | Container registry |

### Observability & Security

| Badge | Technology | Purpose |
|-------|------------|---------|
| <img src="https://img.shields.io/badge/Prometheus-E6522C?style=flat-square&logo=prometheus&logoColor=white" alt="Prometheus"/> | **Prometheus** | Metrics collection |
| <img src="https://img.shields.io/badge/Grafana-F46800?style=flat-square&logo=grafana&logoColor=white" alt="Grafana"/> | **Grafana** | Visualization & dashboards |
| <img src="https://img.shields.io/badge/Trivy-2F77C8?style=flat-square&logo=aqua&logoColor=white" alt="Trivy"/> | **Trivy** | Vulnerability scanning |
| <img src="https://img.shields.io/badge/PostgreSQL-4169E1?style=flat-square&logo=postgresql&logoColor=white" alt="PostgreSQL"/> | **PostgreSQL** | Relational database |

---

## Project Structure

```
.
├── services/                       # Microservices
│   ├── api-gateway/                # API Gateway (routing, metrics)
│   ├── user-service/               # User management CRUD
│   └── order-service/              # Order management CRUD
├── charts/                         # Helm Charts
│   ├── api-gateway/
│   ├── user-service/
│   └── order-service/
├── infra/                          # Infrastructure
│   ├── terraform/                  # AWS (VPC, EKS, RDS, ECR)
│   │   ├── modules/
│   │   │   ├── vpc/
│   │   │   ├── eks/
│   │   │   ├── ecr/
│   │   │   └── rds/
│   │   ├── main.tf
│   │   ├── variables.tf
│   │   └── outputs.tf
│   └── argocd/                     # GitOps App Definitions
│       └── applications/
├── overlays/                       # Kustomize Environments
│   ├── dev/
│   ├── staging/
│   └── prod/
├── monitoring/                     # Observability
│   ├── prometheus/
│   │   ├── prometheus.yml
│   │   └── rules/
│   └── grafana/
│       └── dashboards/
├── scripts/                        # Utilities
├── .github/workflows/              # CI/CD Pipelines
├── docker-compose.yml              # Local Development
└── Makefile                        # Common Commands
```

---

## Quick Start

### Prerequisites

- [Docker](https://docs.docker.com/get-docker/)
- [Make](https://www.gnu.org/software/make/)
- [Go](https://go.dev/doc/install) (for local development)

### Run Locally

```bash
# Clone the repository
git clone https://github.com/Hrishi2861/cicd-platform.git
cd cicd-platform

# Start all services (API, DB, Prometheus, Grafana)
make dev-up

# View logs
docker compose logs -f

# Stop services
make dev-down
```

### API Endpoints

<p align="center">
  <img src="https://img.shields.io/badge/🚪_API_Gateway-localhost:8080-blue?style=for-the-badge" alt="API Gateway"/>
  <img src="https://img.shields.io/badge/👤_User_Service-localhost:8081-green?style=for-the-badge" alt="User Service"/>
  <img src="https://img.shields.io/badge/📦_Order_Service-localhost:8082-orange?style=for-the-badge" alt="Order Service"/>
</p>

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/health` | Health check |
| `GET` | `/metrics` | Prometheus metrics |
| `GET` | `/api/v1/users` | List all users |
| `POST` | `/api/v1/users` | Create a user |
| `GET` | `/api/v1/users/{id}` | Get user by ID |
| `GET` | `/api/v1/orders` | List all orders |
| `POST` | `/api/v1/orders` | Create an order |
| `GET` | `/api/v1/orders/{id}` | Get order by ID |

### Try It Out

```bash
# Create a user
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"email": "john@example.com", "name": "John Doe"}'

# Create an order
curl -X POST http://localhost:8080/api/v1/orders \
  -H "Content-Type: application/json" \
  -d '{"user_id": "<user-id>", "product_name": "Laptop", "quantity": 1, "total_price": 999.99}'

# List users
curl http://localhost:8080/api/v1/users

# List orders
curl http://localhost:8080/api/v1/orders
```

---

## Monitoring Dashboards

<p align="center">
  <img src="https://img.shields.io/badge/📊_Prometheus-localhost:9090-E6522C?style=for-the-badge" alt="Prometheus"/>
  <img src="https://img.shields.io/badge/📈_Grafana-localhost:3000-F46800?style=for-the-badge" alt="Grafana"/>
</p>

| Service | URL | Credentials |
|---------|-----|-------------|
| **Prometheus** | http://localhost:9090 | — |
| **Grafana** | http://localhost:3000 | `admin` / `admin` |

### Alerting Rules

- **High Error Rate** — 5xx responses exceed 5%
- **High Latency** — P95 response time > 1s
- **Pod Restarts** — Container restart detected
- **High CPU** — CPU usage > 80% for 10+ minutes
- **High Memory** — Memory usage > 90% for 10+ minutes

---

## CI/CD Pipeline

### Workflow

```
Pull Request → Test → Lint → Security Scan → Merge
                                           ↓
                                    Build & Push
                                           ↓
                                    Update Manifests
                                           ↓
                                    ArgoCD Auto-Sync
```

### Pipeline Stages

| Stage | Tool | Description |
|-------|------|-------------|
| **Test** | `go test` | Unit tests with coverage reports |
| **Lint** | `golangci-lint` | Static code analysis |
| **Security** | **Trivy** | Container & dependency vulnerability scan |
| **Build** | **Docker Buildx** | Multi-stage builds with caching |
| **Push** | **GHCR** | Push to GitHub Container Registry |
| **Deploy** | **ArgoCD** | GitOps auto-sync to Kubernetes |

### Triggers

| Event | Actions |
|-------|---------|
| **Pull Request** | Test + Lint |
| **Merge to main** | Full pipeline (test → lint → scan → build → push → deploy) |

---

## Infrastructure Deployment

### Prerequisites

```bash
# Required tools
brew install terraform kubectl helm awscli   # macOS
# or use your package manager
```

### Deploy AWS Infrastructure

```bash
cd infra/terraform
terraform init
terraform plan -var="db_username=admin" -var="db_password=securepassword"
terraform apply -var="db_username=admin" -var="db_password=securepassword"
```

### Deploy ArgoCD

```bash
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
kubectl apply -f infra/argocd/applications.yaml
```

### Access ArgoCD UI

```bash
kubectl port-forward svc/argocd-server -n argocd 8080:443
# Username: admin
# Password: kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
```

---

## Environment Configuration

<p align="center">
  <img src="https://img.shields.io/badge/Dev-1_replica_minimal-brightgreen?style=for-the-badge" alt="Dev"/>
  <img src="https://img.shields.io/badge/Staging-2_replicas_standard-yellow?style=for-the-badge" alt="Staging"/>
  <img src="https://img.shields.io/badge/Prod-3_replicas_high-Multi_AZ-red?style=for-the-badge" alt="Prod"/>
</p>

| Environment | Replicas | Resources | Multi-AZ | Deploy Command |
|-------------|----------|-----------|----------|----------------|
| **Dev** | 1 | Minimal | No | `kubectl apply -k overlays/dev/` |
| **Staging** | 2 | Standard | No | `kubectl apply -k overlays/staging/` |
| **Production** | 3 | High | Yes | `kubectl apply -k overlays/prod/` |

---

## Security Features

<p align="center">
  <img src="https://img.shields.io/badge/🔒_Image_Scanning-Trivy-green?style=for-the-badge" alt="Scanning"/>
  <img src="https://img.shields.io/badge/🔐_Encryption-AES256-orange?style=for-the-badge" alt="Encryption"/>
  <img src="https://img.shields.io/badge/🛡️_IAM_Least_Privilege-AWS-blue?style=for-the-badge" alt="IAM"/>
  <img src="https://img.shields.io/badge/🔑_Secrets_Management-K8s_Secrets-red?style=for-the-badge" alt="Secrets"/>
</p>

- **Image Scanning** — Trivy vulnerability scanning in CI pipeline
- **ECR Lifecycle** — Automatic cleanup of old images (keep last 10)
- **IAM Roles** — Least-privilege roles for EKS cluster and nodes
- **Secrets Management** — Kubernetes secrets for database credentials
- **Network Isolation** — VPC security groups restricting RDS access
- **Encryption at Rest** — RDS storage encryption, ECR image encryption

---

## Portfolio Highlights

<p align="center">
  <img src="https://img.shields.io/badge/✅_Microservices-API_Gateway_Pattern-blue?style=for-the-badge" alt="Microservices"/>
  <img src="https://img.shields.io/badge/✅_Infrastructure_as_Code-Terraform-7B42BC?style=for-the-badge" alt="IaC"/>
  <img src="https://img.shields.io/badge/✅_CI/CD_Pipeline-GitHub_Actions-2088FF?style=for-the-badge" alt="CI/CD"/>
</p>
<p align="center">
  <img src="https://img.shields.io/badge/✅_GitOps-ArgoCD_Continuous_Delivery-1e40af?style=for-the-badge" alt="GitOps"/>
  <img src="https://img.shields.io/badge/✅_Container_Orchestration-Kubernetes-326CE5?style=for-the-badge" alt="K8s"/>
  <img src="https://img.shields.io/badge/✅_Observability-Prometheus_&_Grafana-E6522C?style=for-the-badge" alt="Observability"/>
</p>
<p align="center">
  <img src="https://img.shields.io/badge/✅_Security_Image_Scanning-Trivy-2F77C8?style=for-the-badge" alt="Security"/>
  <img src="https://img.shields.io/badge/✅_Multi_Environment-Dev_Staging_Prod-yellow?style=for-the-badge" alt="Environments"/>
  <img src="https://img.shields.io/badge/✅_Auto_Scaling-HPA_CPU_Based-green?style=for-the-badge" alt="HPA"/>
</p>
<p align="center">
  <img src="https://img.shields.io/badge/✅_Production_Ready-Health_Checks_&_Limits-red?style=for-the-badge" alt="Production"/>
  <img src="https://img.shields.io/badge/✅_Package_Management-Helm_&_Kustomize-0F1689?style=for-the-badge" alt="Helm"/>
</p>

---

<p align="center">
  <sub>Built with ❤️ using Go, Docker, Kubernetes, and Terraform By Hrishikesh</sub>
</p>
