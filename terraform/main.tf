provider "aws" {
  region = var.region
  access_key = var.access_key
  secret_key = var.secret_key
}

# Filter out local zones, which are not currently supported
# with managed node groups
data "aws_availability_zones" "available" {
  filter {
    name   = "opt-in-status"
    values = ["opt-in-not-required"]
  }
}

locals {
  cluster_name = "go-kafka-order-producer-eks"
}


module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "5.8.1"

  name = "go-kafka-order-producer-vpc"

  cidr = "10.0.0.0/16"
  azs  = slice(data.aws_availability_zones.available.names, 0, 3)

  private_subnets = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
  public_subnets  = ["10.0.4.0/24", "10.0.5.0/24", "10.0.6.0/24"]

  enable_nat_gateway   = true
  single_nat_gateway   = true
  enable_dns_hostnames = true

  public_subnet_tags = {
    "kubernetes.io/role/elb" = 1
  }

  private_subnet_tags = {
    "kubernetes.io/role/internal-elb" = 1
  }
}

module "eks" {
  source  = "terraform-aws-modules/eks/aws"
  version = "20.8.5"

  cluster_name    = local.cluster_name
  cluster_version = "1.31"

  cluster_endpoint_public_access           = true
  enable_cluster_creator_admin_permissions = true

  vpc_id     = module.vpc.vpc_id
  subnet_ids = module.vpc.private_subnets

  eks_managed_node_group_defaults = {
    ami_type = "AL2_x86_64"

  }

  eks_managed_node_groups = {
    one = {
      name = "node-group-1"

      instance_types = ["t3.small"]

      min_size     = 1
      max_size     = 3
      desired_size = 2
    }

    two = {
      name = "node-group-2"

      instance_types = ["t3.small"]

      min_size     = 1
      max_size     = 2
      desired_size = 1
    }
  }
}


resource "aws_iam_user" "github" {
  name = "GithubCICD"
}


data "aws_iam_policy_document" "github_cicd_policy" {
  statement {
    actions   = ["ecr:GetAuthorizationToken",
    "ecr:BatchGetImage",
    "ecr:GetDownloadUrlForLayer",
    "ecr:DescribeImages",
    "ecr:InitiateLayerUpload",
    "ecr:PutImage",
    "ecr:UploadLayerPart",
    "ecr:CompleteLayerUpload",
    "ecr:BatchCheckLayerAvailability"]
    resources = ["*"]
    effect = "Allow"
  }
  statement {
    actions   = ["eks:DescribeCluster","eks:ListClusters","eks:UpdateClusterConfig","eks:*"]
    resources = ["*"]
    effect = "Allow"
  }
}

resource "aws_iam_policy" "policy" {
  name        = "github_cicd-policy"
  description = "Policy to github cicd"
  policy = data.aws_iam_policy_document.github_cicd_policy.json
}

resource "aws_iam_user_policy_attachment" "attachment" {
  user       = aws_iam_user.github.name
  policy_arn = aws_iam_policy.policy.arn
}

resource "awscc_eks_access_entry" "githubCICD" {
  cluster_name    = local.cluster_name
  principal_arn     = aws_iam_user.github.arn
  type              = "STANDARD"
  access_policies = [
    {
      access_scope = {
        type       = "namespace"
        namespaces = ["default"]
      }
      policy_arn = "arn:aws:eks::aws:cluster-access-policy/AmazonEKSAdminPolicy"
    }
  ]
}