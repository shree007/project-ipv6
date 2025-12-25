
terraform {
  source = "../terraform-modules/vpc"
}

include "root" {
  path = find_in_parent_folders()
}

inputs = {
  aws_region                 = "us-east-1"
  vpc_cidr                  = "10.0.0.0/16"
  vpc_name                  = "multi-az-vpc"
  igw_name                  = "multi-az-igw"
  public_rt_name            = "multi-az-public-rt"
  public_subnet_name_prefix = "multi-az-public-subnet"
  private_subnet_name_prefix = "multi-az-private-subnet"
}

remote_state {
  backend = "s3"
  config = {
    bucket         = "my-terraform-state-bucket"
    key            = "multi-az-vpc/terraform.tfstate"
    region         = "us-east-1"
    encrypt        = true
    dynamodb_table = "terraform-lock"
  }
}
