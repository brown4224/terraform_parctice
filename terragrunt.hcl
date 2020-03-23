terraform {
  extra_arguments "common_vars" {
    commands = ["plan", "apply"]

    arguments = [
      "-var-file=${get_parent_terragrunt_dir()}/vars/vpc.tfvars"
    ]
  }
}

locals {
  # Automatically load region-level variables
  region_vars = read_terragrunt_config(find_in_parent_folders("region.hcl"))
  aws_region   = local.region_vars.locals.aws_region
}

inputs = merge(
  local.region_vars.locals,
)

# generate "provider" {
#   path = "provider.tf"
#   if_exists = "overwrite_terragrunt"
#   contents = <<EOF
# provider "aws" {
#   profile = "default"
#   region = "${local.aws_region}"

# }
# EOF
# }
