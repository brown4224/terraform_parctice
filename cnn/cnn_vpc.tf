resource "aws_vpc" "vpc" {
  cidr_block = "10.200.0.0/16"
  tags = {
      Name = "terraform test"
  }
}

resource "aws_security_group" "open_ports" {
    name = "cnn_test"
    description = "This is only a test"
    vpc_id = "${aws_vpc.vpc.id}"

    depends_on = ["aws_vpc.vpc"]
    dynamic "ingress" {
        for_each = var.ports
        content {
            from_port = ingress.value
            to_port = ingress.value
            protocol = "tcp"
            cidr_blocks = var.cidr
        }
    }
    ingress {
        from_port = "-1"
        to_port = "-1"
        protocol = "icmp"
        cidr_blocks = var.cidr
    }
    egress {
        from_port = "0"
        to_port = "0"
        protocol = "-1"
        cidr_blocks = ["0.0.0.0/0"]
    }
    tags = {
        Name = "CNN Test"
    }
}

resource "aws_internet_gateway" "gw" {
    vpc_id = "${aws_vpc.vpc.id}"
    tags = {
        Name = "terraform-gw"
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
  cidr_block = "10.200.1.0/24"

  tags = {
    Name = "Terraform-Subnet"
  }
}
resource "aws_route_table_association" "a" {
    subnet_id = "${aws_subnet.main.id}"
    route_table_id = "${data.aws_route_table.selected.id}"
  
}
