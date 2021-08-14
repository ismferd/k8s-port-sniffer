variable "bucket_name" {
    default = "node-port-scanner-default"
}

variable "acl_value" {
    default = "private"
}

variable "aws_region" {
    default = "us-east-2"
}

variable "aws_access_key" {
  type = string
  default = "test"
}

variable "aws_secret_key" {
  type = string
  default = "test"
}

variable "versioning" {
    default = true
}

variable "s3_endpoint" {
    default =   "https://s3.us-central-2.amazonaws.com"
}
variable "sts_endpoint" {
    default =   "https://sts.us-central-2.amazonaws.com"
}
variable "skip_credentials_validation" {
    default =   false
}

variable "skip_metadata_api_check" {
    default =   false
}
variable "skip_requesting_account_id" {
    default =   false
}
