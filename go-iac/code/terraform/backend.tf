terraform {
  required_version = ">= 0.12"

  backend "s3" {
    bucket = "my-infra"
    key    = "states/envs"
    region = "us-east-1"
  }
}