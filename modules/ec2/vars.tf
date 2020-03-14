variable "name" {
  description = "The friendly name"
  type        = string
}

variable "ami" {
  description = "The ami"
  type        = string
}

variable "instance" {
  description = "The instance type"
  type        = string
}

variable "keypair" {
  description = "The keypair to use"
  type        = string
}

variable "subnet_id" {
  description = "The subnet id for the instance"
  type        = string
}

variable "security_group_ids" {
  description = "The security group id for the instance"
  type        = list(string)
}