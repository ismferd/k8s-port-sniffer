resource "aws_s3_bucket" "bucket" {
   bucket = var.bucket_name
   acl = var.acl_value
   versioning {
      enabled = var.versioning
   }
   tags = {
     Name = var.bucket_name
     Environment = "dev"
     Owner = "sre"
     Managed = "terraform"
   }
}