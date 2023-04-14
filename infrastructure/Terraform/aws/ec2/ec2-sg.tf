 resource "aws_security_group" "vpc-ssh" {
   name = "vpc-ssh"
   description = "Dev VPC SSH"
   ingress {
     cidr_blocks = [ "0.0.0.0/0" ]
     description = "Alow port 22"
     from_port = 22
     protocol = "tcp"
     to_port = 22
   } 

  egress {
    cidr_blocks = [ "0.0.0.0/0" ]
    description = "Allow all outbund to all ip and ports"
    from_port = 0
    protocol = "-1"
    to_port = 0
  } 

  

  tags = {
    "Name" = "vpc-ssh"
  }

 }