# terraform {
#   backend "s3" {
#     bucket         = "opolis-terraform-state-prod"
#     key            = "pre-prod-infra.tfstate"
#     region         = "us-east-1"
#     dynamodb_table = "opolis-tf-state-lock-pre-prod"
#   }
# }

provider "aws" {
  region = "us-east-1"
}