provider "aws" {
  profile = "default"
  region = "us-east-1"
}

variable "aws_region" {
  default = "us-east-1"
}
variable "aws_account_number" {
  default = "728160576949"
}
variable "aws_resource_prefix" {
  default = "certbot-"
}
variable "aws_task_definition" {
  default = "go-gen-ssl"
}