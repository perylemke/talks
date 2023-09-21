terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.5.0"
    }
  }
}

# Configure the AWS Provider
provider "aws" {
  default_tags {
    tags = {
      environment = terraform.workspace
      owner       = "pery"
      created_by  = "pery"
    }
  }
}