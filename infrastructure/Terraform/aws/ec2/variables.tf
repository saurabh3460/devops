variable "aws_region" {
  type = string
  default = "us-east-1"
}

variable "instance_type" {
  type = string
  default = "t2.micro"
}

variable "instance_keypair" {
  description = "AWS EC2 Key Pair that need to be associated with EC2 Instance"
  type = string
  default = "terraform-key"
}


# AWS EC2 Instance Type - List
variable "instance_type_list" {
  description = "EC2 Instance Type"
  type = set(string)
  default = ["t2.micro", "t3.small", "t3.large", "t3.large"]  
}

# AWS EC2 Instance Type - Map
variable "instance_type_map" {
  description = "EC2 Instance Type"
  type = map(string)
  default = {
    "dev" = "t2.micro"
    "qa" = "t3.small"
    "prod" = "t3.large"
  }
}

variable "ingress" {
  description = "EC2 Instance Type"
  type = object({cidr=string, port=number })
  default = {
    cidr = "0.0.0.0/0"
    port = 22
  }
}

