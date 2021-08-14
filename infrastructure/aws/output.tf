output "bucket_name" {
  value = aws_s3_bucket.bucket.bucket
}

output "acl" {
  value = aws_s3_bucket.bucket.acl
}

output "tags" {
  value = aws_s3_bucket.bucket.tags
}