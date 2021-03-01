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
    default = "vpc-0c616a64"
}