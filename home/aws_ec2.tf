# # Sean McGlincy Home Files

# resource "aws_instance" "ec2" {
#   ami           = "ami-027a14492d667b8f5"
#   instance_type = "t2.micro"
#   vpc_security_group_ids = [aws_security_group.sg.id]
#   subnet_id = "${aws_subnet.main.id}"
#   associate_public_ip_address = true
#   key_name = "dev-ansable-key-01"
#   tags = {
#       Name = "Terraform-Windows"
#   }

#   provisioner "remote-exec" {
#     inline = [
#       "$url = 'https://raw.githubusercontent.com/ansible/ansible/devel/examples/scripts/ConfigureRemotingForAnsible.ps1'",
#       "$file = '$env:temp\\ConfigureRemotingForAnsible.ps1'",
#       "(New-Object -TypeName System.Net.WebClient).DownloadFile($url, $file)",
#       "powershell.exe -ExecutionPolicy ByPass -File $file"
#     ]
#   }
# }
