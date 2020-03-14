include {
  path = find_in_parent_folders()
}

terraform {
  source = "../../modules/vpc"
}

# These are the variables we have to pass in to use the module specified in the terragrunt configuration above
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