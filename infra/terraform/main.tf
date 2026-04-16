terraform {
  required_version = ">= 1.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }

  backend "s3" {
    bucket         = "terraform-state-cicd-platform"
    key            = "infrastructure/terraform.tfstate"
    region         = "us-east-1"
    dynamodb_table = "terraform-lock-cicd-platform"
    encrypt        = true
  }
}

provider "aws" {
  region = var.aws_region

  default_tags {
    tags = {
      ManagedBy   = "terraform"
      Project     = "cicd-platform"
      Environment = var.environment
    }
  }
}

module "vpc" {
  source = "./modules/vpc"

  vpc_cidr     = var.vpc_cidr
  aws_region   = var.aws_region
  environment  = var.environment
  project_name = var.project_name
}

module "eks" {
  source = "./modules/eks"

  cluster_name    = var.cluster_name
  cluster_version = var.cluster_version
  vpc_id          = module.vpc.vpc_id
  subnet_ids      = module.vpc.private_subnet_ids
  node_groups     = var.node_groups
  environment     = var.environment
}

module "ecr" {
  source = "./modules/ecr"

  repositories = var.ecr_repositories
  environment  = var.environment
}

module "rds" {
  source = "./modules/rds"

  db_username       = var.db_username
  db_password       = var.db_password
  subnet_ids        = module.vpc.private_subnet_ids
  security_group_id = module.vpc.db_security_group_id
  environment       = var.environment
}
