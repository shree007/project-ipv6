variable "region" {
  description = "The AWS region to deploy EKS cluster"
  type        = string
  default     = "us-west-2"
}

variable "cluster_name" {
  description = "The name of the EKS cluster"
  type        = string
  default     = "my-eks-cluster"
}

variable "node_group_name" {
  description = "The name of the EKS node group"
  type        = string
  default     = "eks-node-group"
}

variable "node_instance_type" {
  description = "The EC2 instance type for the nodes"
  type        = string
  default     = "t3.medium"
}

variable "desired_capacity" {
  description = "The desired number of worker nodes"
  type        = number
  default     = 2
}

variable "min_size" {
  description = "The minimum number of worker nodes"
  type        = number
  default     = 1
}

variable "max_size" {
  description = "The maximum number of worker nodes"
  type        = number
  default     = 3
}

variable "vpc_id" {
  description = "The ID of the VPC"
  type        = string
}

variable "subnet_ids" {
  description = "The IDs of the subnets"
  type        = list(string)
}

variable "cluster_ip_family" {
  description = "Kubernetes cluster IP family (ipv4 or ipv6)"
  type        = string
  default     = "ipv6"
}

variable "create_cni_ipv6_iam_policy" {
  description = "Attach the IPv6 CNI IAM policy to node roles"
  type        = bool
  default     = true
}

variable "cluster_addons" {
  description = "EKS cluster addons configuration"
  type        = map(any)
  default = {
    vpc-cni = {
      most_recent = true
    }
  }
}
