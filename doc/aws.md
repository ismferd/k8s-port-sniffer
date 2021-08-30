# AWS 

In order to store information collected by k8s-port-scanner of ports opened we need a s3 bucket.

## Terraform

You can find the terraform code [here](infrastructure/aws).

### Variables

* bucket_name: The name of the bucket, default: node-port-scanner-default.
* acl_value: Access to  your buckets, default: private.
* aws_region: Region where the AWS provider will be initialized, default: us-east-2.
* aws_access_key: Your AWS_ACCESS_KEY, default: test.
* aws_secret_key: Your AWS_SECRET_KEY, default: test.
* versioning: Versioning of the bucket, default: true.
* s3_endpoint: The endpoint to s3, useful to use localstack, default: https://s3.us-east-2.amazonaws.com.
* skip_credentials_validation: skip credentials validation, useful to use terratest, default = false.
* skip_metadata_api_check: skip metadata api check, useful to use terratest, default = false.
* skip_requesting_account_id: skip requesting account id, useful to use terratest, default = false.

### Deployment

Running the command `make deploy_aws` you will need set both env vars:
* TF_VAR_aws_access_key
* TF_VAR_aws_secret_key

These both env vars will arrive to terraform to validate your credentials agains AWS.

`TF_VAR_aws_access_key=MY_AWS_KEY TF_VAR_aws_secret_key=MY_AWS_SECRET make deploy_aws` will do:
- Up localstack.
- Test infrastructure with terratest.
- Down localstack.
- Terraform apply.