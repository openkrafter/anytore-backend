terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
      #      version = "~> 4.16"
    }
  }

  required_version = ">= 1.2.0"

  backend "s3" {
    region = "ap-northeast-1"
  }
}

provider "aws" {
  region = "ap-northeast-1"
}

resource "aws_apprunner_service" "anytore-backend" {
  service_name = var.service_name

  source_configuration {
    image_repository {
      image_configuration {
        port = "80"
      }
      image_identifier      = format("%s.dkr.ecr.ap-northeast-1.amazonaws.com/%s:%s", var.account_id, var.repository_name, var.image_tag)
      image_repository_type = "ECR"
    }
    authentication_configuration {
      access_role_arn = format("arn:aws:iam::%s:role/service-role/AppRunnerECRAccessRole", var.account_id)
    }
    auto_deployments_enabled = true
  }

  instance_configuration {
    cpu               = "1024"
    instance_role_arn = format("arn:aws:iam::%s:role/AppRunnerInstanceRole", var.account_id)
    memory            = "2048"
  }

  health_check_configuration {
    healthy_threshold   = "1"
    interval            = "10"
    path                = "/"
    protocol            = "TCP"
    timeout             = "5"
    unhealthy_threshold = "5"
  }

  auto_scaling_configuration_arn = format("arn:aws:apprunner:ap-northeast-1:%s:autoscalingconfiguration/DefaultConfiguration/1/00000000000000000000000000000001", var.account_id)

  tags = {
    Name = "anytore"
  }
}
