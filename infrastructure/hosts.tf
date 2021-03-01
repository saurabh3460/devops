terraform {
    required_providers {
        aws = {
            source = "hashicorp/aws"
            version = "~>2.7.3"
        }
    }
}

provider "aws" {
    profile = "default"
    region = "ap-south-1"

}

resource "aws_key_pair" "ansible"{
    key_name = var.host01.key_name
    public_key = file("${abspath(path.cwd)}/ansible.pub")
}

resource "aws_instance" "host01" {
    ami = var.host01["ami"]
    instance_type = var.host01["instance_type"]
    key_name = var.host01.key_name
}


