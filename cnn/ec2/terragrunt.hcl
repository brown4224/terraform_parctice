include {
  path = find_in_parent_folders()
}

dependency "vpc" {
  config_path = "../vpc"
}

terraform {
  source = "../../modules/ec2"
}

inputs = {
    name = "Terraform-Windows"
    ami = "ami-0b940a1059f928462"
    instance = "t2.micro"
    keypair = "dev-ansable-key-01"
    subnet_id = dependency.vpc.outputs.subnet_id
    security_group_ids = [dependency.vpc.outputs.security_group_id]
}
