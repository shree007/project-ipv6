output "cluster_id" {
  description = "The ID of the EKS cluster"
  value       = module.eks.cluster_id
}

output "cluster_endpoint" {
  description = "The endpoint of the EKS cluster"
  value       = module.eks.cluster_endpoint
}

output "cluster_arn" {
  description = "The ARN of the EKS cluster"
  value       = module.eks.cluster_arn
}

output "kubeconfig" {
  description = "Kubeconfig file contents for the cluster"
  value       = module.eks.kubeconfig
}

output "node_group_names" {
  description = "Names of the node groups created"
  value       = module.eks.node_groups
}