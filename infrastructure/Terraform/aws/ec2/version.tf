terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 3.63.0"
    }
  }
}

provider "aws" {
  region = "us-east-1"
  alias  = "india"
}

# provider "aws" {
#   region = "us-west-2"
#   alias  = "oregon"
# }