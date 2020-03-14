resource "aws_instance" "ec2"{
    ami = "ami-09d069a04349dc3cb"
    instance_type = "t2.micro"
    vpc_security_group_ids = [aws_security_group.open_ports.id]
    subnet_id =" ${aws_subnet.main.id}"
    associate_public_ip_address = true
    key_name = "dev-ansable-key-01"
    tags = {
        Name = "Terraform-Linux"
    }
}
