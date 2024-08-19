variable "instance_name" {
  description = "Value of the Name tag for the EC2 instance"
  type        = string
  default     = "orto-server"
}

variable "bucket_name" {
  description = "Value for the name of the s3 bucket"
  type        = string
  default     = "orto-deployment-assets"
}

variable "code_deploy_app" {
  description = "Value of the code deploy app"
  type        = string
  default     = "ormanel"
}

variable "deployment_group_name" {
  description = "Value of the code deploy deployment group name"
  type        = string
  default     = "orto"
}


