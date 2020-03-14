# # Sean McGlincy Home Files

# variable "cidr" {
#   default = ["136.55.41.30/32"]
# }
# variable "ports" {
#   type = map(string)
#   default = {
#     ssh = "22"
#     rdp = "3389"
#   }
# }

# provider "aws" {
#     profile = "default"
#     region = "us-east-1"
# }

# resource "aws_vpc" "main" {
#     cidr_block = "10.1.0.0/16"
#     tags = {
#         Name="Terraform-Test"
#     }
# }

# resource "aws_security_group" "sg" {
#   name = "Terraform-bastion"
#   description = "Terraform Test"
#   vpc_id = "${aws_vpc.main.id}"
#   depends_on = ["aws_vpc.main"]
#   dynamic "ingress" {
#     for_each = var.ports
#       content {
#           from_port = ingress.value
#           to_port = ingress.value
#           protocol = "tcp"
#           cidr_blocks = var.cidr
#       }
#   }

#   ingress {
#     from_port       = "-1"
#     to_port         = "-1"
#     protocol = "icmp"
#     cidr_blocks = var.cidr
#   }
#   egress {
#     from_port       = "0"
#     to_port         = "0"
#     protocol        = "-1"
#     cidr_blocks     = ["0.0.0.0/0"]
#   }
#   tags = {
#     Name = "Terraform-Test"
#   }
# }

# resource "aws_internet_gateway" "gw" {
#   vpc_id = "${aws_vpc.main.id}"
#   tags = {
#     Name = "terraform-gw"
#   }
# }

# data "aws_route_table" "selected" {
#   vpc_id = "${aws_vpc.main.id}"
# }

# resource "aws_route" "route" {
#   route_table_id            = "${data.aws_route_table.selected.id}"
#   destination_cidr_block    = "0.0.0.0/0"
#   gateway_id = "${aws_internet_gateway.gw.id}"
# }

# resource "aws_subnet" "main" {
#   vpc_id     = "${aws_vpc.main.id}"
#   cidr_block = "10.1.0.0/24"

#   tags = {
#     Name = "Terraform-Subnet"
#   }
# }

# resource "aws_route_table_association" "a" {
#   subnet_id      = "${aws_subnet.main.id}"
#   route_table_id = "${data.aws_route_table.selected.id}"
# }
