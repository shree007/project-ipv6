variable "aws_region" {
  description = "The AWS region to deploy resources in"
  type        = string
  default     = "us-east-1"
}

variable "vpc_cidr" {
  description = "The CIDR block for the VPC"
  type        = string
  default     = "172.20.0.0/16"
}

variable "ipv6_cidr" {
  description = "The IPv6 CIDR block for the VPC"
  type        = string
  default     = "2600:1f18:abcd::/56"
}

variable "vpc_name" {
  description = "The name of the VPC"
  type        = string
  default     = "project-ipv6"
}

variable "igw_name" {
  description = "The name of the Internet Gateway"
  type        = string
  default     = "multi-az-igw"
}

variable "public_rt_name" {
  description = "The name of the public route table"
  type        = string
  default     = "multi-az-public-rt"
}

variable "public_subnet_name_prefix" {
  description = "Prefix for public subnet names"
  type        = string
  default     = "multi-az-public-subnet"
}

variable "private_subnet_name_prefix" {
  description = "Prefix for private subnet names"
  type        = string
  default     = "multi-az-private-subnet"
}
