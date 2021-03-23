variable "host01"  {
    type  = map(string)
    default = {
        ami = "ami-073c8c0760395aab8"
        instance_type = "t2.micro"
        key_name = "ansible"
    }
}

variable "vpc_id" {
    type = string
    default = "vpc-dc956bb7"
}

variable "inventory_path" {
    type = string
    default = "../playbooks/inventory"
}

variable "private_key_path" {
    type = string
    default = "../infrastructure/ansible"
}