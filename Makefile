.PHONY: help dev-up dev-down test lint build deploy

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

dev-up: ## Start all services locally
	docker compose up -d --build

dev-down: ## Stop all services
	docker compose down

test: ## Run tests for all services
	@echo "Running tests for api-gateway..."
	@cd services/api-gateway && go test ./... -v
	@echo "Running tests for user-service..."
	@cd services/user-service && go test ./... -v
	@echo "Running tests for order-service..."
	@cd services/order-service && go test ./... -v

lint: ## Lint all services
	@echo "Linting api-gateway..."
	@cd services/api-gateway && golangci-lint run || true
	@echo "Linting user-service..."
	@cd services/user-service && golangci-lint run || true
	@echo "Linting order-service..."
	@cd services/order-service && golangci-lint run || true

build: ## Build all Docker images
	docker compose build

deploy-infra: ## Deploy infrastructure with Terraform
	cd infra/terraform && terraform init && terraform plan && terraform apply -auto-approve

deploy-apps: ## Deploy applications with ArgoCD
	kubectl apply -f infra/argocd/namespace.yaml
	kubectl apply -f infra/argocd/install.yaml
	kubectl apply -f infra/argocd/applications.yaml

destroy-infra: ## Destroy infrastructure
	cd infra/terraform && terraform destroy -auto-approve

screenshot-argocd: ## Spin up local kind + ArgoCD + api-gateway for a screenshot
	./scripts/argocd-screenshot.sh

screenshot-argocd-destroy: ## Tear down the kind cluster used for the screenshot
	./scripts/argocd-screenshot.sh --destroy
