provider "aws" {
  region = var.aws_region
}

resource "aws_instance" "ec2" {
  ami           = var.ami
  instance_type = var.instance
  vpc_security_group_ids = var.security_group_ids
  subnet_id = var.subnet_id
  associate_public_ip_address = true
  key_name = var.keypair
  tags = {
      Name = var.name
  }

  # provisioner "remote-exec" {
  #   inline = [
  #     "$url = 'https://raw.githubusercontent.com/ansible/ansible/devel/examples/scripts/ConfigureRemotingForAnsible.ps1'",
  #     "$file = '$env:temp\\ConfigureRemotingForAnsible.ps1'",
  #     "(New-Object -TypeName System.Net.WebClient).DownloadFile($url, $file)",
  #     "powershell.exe -ExecutionPolicy ByPass -File $file"
  #   ]
  # }
}
