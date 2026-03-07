# Cloud-Native CI/CD Platform

Production-ready microservices platform with full CI/CD, GitOps & observability.

> Work in progress — building this out incrementally as a portfolio project.

## Goals

- Microservices in Go behind an API gateway
- Containerized local dev with Docker Compose
- Infrastructure as Code on AWS with Terraform (VPC, EKS, ECR, RDS)
- Kubernetes packaging with Helm + Kustomize overlays (dev/staging/prod)
- GitOps delivery with ArgoCD
- CI/CD with GitHub Actions (test, lint, security scan, build)
- Observability with Prometheus + Grafana

## Services

| Service | Port | Description |
|---------|------|-------------|
| api-gateway | 8080 | Request routing and metrics |
| user-service | 8081 | User management |
| order-service | 8082 | Order management |

## Status

Early scaffold. See commit history for progress.

## License

MIT
