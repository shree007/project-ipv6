module "eks" {
  source          = "terraform-aws-modules/eks/aws"
  cluster_name    = var.cluster_name
  cluster_version = "1.30"
  subnets         = var.subnet_ids
  vpc_id          = var.vpc_id
  cluster_ip_family             = var.cluster_ip_family
  create_cni_ipv6_iam_policy     = var.create_cni_ipv6_iam_policy
  cluster_addons                 = var.cluster_addons

  node_groups = {
    default = {
      desired_capacity = var.desired_capacity
      max_capacity     = var.max_size
      min_capacity     = var.min_size

      instance_type = var.node_instance_type
    }
  }
}
