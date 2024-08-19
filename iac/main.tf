provider "aws" {
  region = "us-east-1"
}

# Ec2 Instance Provisioning
resource "aws_instance" "app_server" {
  ami           = "ami-04a81a99f5ec58529"
  instance_type = "t2.micro"

  tags = {
    Name = var.instance_name
  }
  user_data            = file("./_setup.sh")
  security_groups      = [aws_default_security_group.app_sg.name]
  iam_instance_profile = aws_iam_instance_profile.ec2_instance_profile.name
}

# S3 Bucket Provisioning
resource "aws_s3_bucket" "app_storage" {
  bucket        = var.bucket_name
  force_destroy = true

  tags = {
    Name        = "orto-bucket"
    Environment = "Dev"
  }
}

# Code Deploy App Provisioning
resource "aws_codedeploy_app" "app_deploy" {
  name = var.code_deploy_app
}

# Code Deploy Deployment Configuration
resource "aws_codedeploy_deployment_config" "app_config" {
  deployment_config_name = "ortodeploy-config"

  minimum_healthy_hosts {
    type  = "HOST_COUNT"
    value = 1
  }
}

# Code Deploy Deployment Group Provisioning
resource "aws_codedeploy_deployment_group" "app_group" {
  app_name               = aws_codedeploy_app.app_deploy.name
  deployment_group_name  = var.deployment_group_name
  service_role_arn       = aws_iam_role.orto_codedeploy_role.arn
  deployment_config_name = aws_codedeploy_deployment_config.app_config.id

  ec2_tag_filter {
    key   = "Name"
    type  = "KEY_AND_VALUE"
    value = var.instance_name
  }

  auto_rollback_configuration {
    enabled = true
    events  = ["DEPLOYMENT_FAILURE"]
  }
}

# Code Deploy Policy
data "aws_iam_policy_document" "code_deploy_assume_role" {
  statement {
    effect = "Allow"

    principals {
      type        = "Service"
      identifiers = ["codedeploy.amazonaws.com"]
    }

    actions = ["sts:AssumeRole"]
  }
}

# IAM Role For CodeDeploy
resource "aws_iam_role" "orto_codedeploy_role" {
  name               = "orto_codedeploy_roles"
  assume_role_policy = data.aws_iam_policy_document.code_deploy_assume_role.json
}


# Attach Policy to IAM Role
resource "aws_iam_role_policy_attachment" "code_deploy_role" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSCodeDeployRole"
  role       = aws_iam_role.orto_codedeploy_role.name
}


# EC2 Policy
data "aws_iam_policy_document" "ec2_assume_role" {
  statement {
    effect = "Allow"

    principals {
      type        = "Service"
      identifiers = ["ec2.amazonaws.com"]
    }

    actions = ["sts:AssumeRole"]
  }
}

# EC2 IAM Role
resource "aws_iam_role" "ec2_role" {
  name               = "ec2_role"
  assume_role_policy = data.aws_iam_policy_document.ec2_assume_role.json
}

# Attaching Policy to Role
resource "aws_iam_role_policy_attachment" "ec2_policy" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess"
  role       = aws_iam_role.ec2_role.name
}


# Creating an Instance Profile for the server
resource "aws_iam_instance_profile" "ec2_instance_profile" {
  name = "ec2_instance_profile"
  role = aws_iam_role.ec2_role.name
}

# AWS VPC
resource "aws_vpc" "mainvpc" {
  cidr_block = "10.1.0.0/16"
}

# Security Group for the app
resource "aws_default_security_group" "app_sg" {
  vpc_id = aws_vpc.mainvpc.id

  ingress {
    protocol  = -1
    self      = true
    from_port = 0
    to_port   = 0
    #cidr_blocks = ["0.0.0.0/0"]
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  tags = {
    Name = "app_security_group"
  }
}


