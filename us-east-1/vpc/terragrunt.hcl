include {
  path = find_in_parent_folders()
}

terraform {
  source = "../../modules/vpc"
}

inputs = {
  name       = "terragrunt_test"
  vpc_cidr = "10.1.0.0/16"
  subnet_cidr = "10.1.0.0/24"
  ingress_cidr =  ["136.55.41.30/32"]
  ports = {
      ssh = "22"
      rdp = "3389"
  }
}