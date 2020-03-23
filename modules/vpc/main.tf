# Meta Requirements for Terraform
# terraform {
#   required_version = "= 0.12.12"
#     required_providers {
#     aws = {
#       version = "= 2.33.0"
#     }
#   }
# }


provider "aws" {
  region = var.aws_region
}

resource "aws_vpc" "vpc" {
  cidr_block = var.vpc_cidr
  tags = {
      Name = format("%s_vpc", var.name)
  }
}

resource "aws_security_group" "open_ports" {
    name = var.name
    description = "This was created by terraform"
    vpc_id = "${aws_vpc.vpc.id}"

    depends_on = ["aws_vpc.vpc"]
    dynamic "ingress" {
        for_each = var.ports
        content {
            from_port = ingress.value
            to_port = ingress.value
            protocol = "tcp"
            cidr_blocks = var.ingress_cidr
        }
    }
    ingress {
        from_port = "-1"
        to_port = "-1"
        protocol = "icmp"
        cidr_blocks = var.ingress_cidr
    }
    egress {
        from_port = "0"
        to_port = "0"
        protocol = "-1"
        cidr_blocks = ["0.0.0.0/0"]
    }
    tags = {
        Name = format("%s_sg", var.name)
    }
}

resource "aws_internet_gateway" "gw" {
    vpc_id = "${aws_vpc.vpc.id}"
    tags = {
        Name = format("%s_gw", var.name)
    }
}

data "aws_route_table" "selected" {
    vpc_id = "${aws_vpc.vpc.id}"
}

resource "aws_route" "route" {
    route_table_id = "${data.aws_route_table.selected.id}"
    destination_cidr_block = "0.0.0.0/0"
    gateway_id = "${aws_internet_gateway.gw.id}"
}

resource "aws_subnet" "main" {
  vpc_id     = "${aws_vpc.vpc.id}"
  cidr_block = var.subnet_cidr

  tags = {
    Name = format("%s_subnet", var.name)
  }
}
resource "aws_route_table_association" "a" {
    subnet_id = "${aws_subnet.main.id}"
    route_table_id = "${data.aws_route_table.selected.id}"
  
}
