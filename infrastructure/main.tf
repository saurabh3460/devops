terraform {
    required_providers {
        aws = {
            source = "hashicorp/aws"
            version = "3.30.0"
        }
    }
}

provider "aws" {
    profile = "default"
    region = "ap-south-1"
}

resource "aws_security_group" "allow_ssh" {
    name = "ansible"
    vpc_id = var.vpc_id

    ingress {
        from_port = 22
        to_port = 22
        protocol = "tcp"
        cidr_blocks = ["0.0.0.0/0"]
    }

    egress {
        from_port = 0
        to_port = 0
        protocol = "-1"
        cidr_blocks = ["0.0.0.0/0"]
    }
}

resource "aws_key_pair" "ansible" {
    key_name = var.host01.key_name
    public_key = file("${abspath(path.cwd)}/ansible.pub")
}

resource "aws_instance" "host01" {
    ami = var.host01["ami"]
    instance_type = var.host01["instance_type"]
    key_name = var.host01.key_name
    security_groups = [aws_security_group.allow_ssh.name]
}

resource "local_file" "host_ip" {
    content  = aws_instance.host01.public_ip
    filename = var.inventory_path
}
