variable "aws_region" {
  description = "The AWS region to deploy to (e.g. us-east-1)"
  type        = string
}

variable "name" {
  description = "The friendly name"
  type        = string
}

variable "vpc_cidr" {
  description = "The address of  the vpc cidr"
  type        = string
}

variable "subnet_cidr" {
  description = "The address of  the vpc cidr"
  type        = string
}

variable "ingress_cidr" {
  description = "The address of the ingress cidr. ie your IP"
  type        = list(string)
}

variable "ports" {
  description = "The address of  the vpc cidr"
  type        = map(string)
  default     = {
				ssh = "22"
				rdp = "3389",
			}
}