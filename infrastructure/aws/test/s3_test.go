package infratest

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// Terraform tests
func TestTerraformAwsS3Example(t *testing.T) {
	t.Parallel()

	// Give this S3 Bucket a unique ID for a name tag so we can distinguish it from any other Buckets provisioned
	// in your AWS account
	expectedName := fmt.Sprintf("node-port-scanner-default-%s", strings.ToLower(random.UniqueId()))
	expectedAcl := "private"

	// Construct the terraform options with default retryable errors to handle the most common retryable errors in
	// terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"bucket_name": expectedName,
		},
	})

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	var terraformOpts = &terraform.Options{
		TerraformDir: "../",
	}
	outputBucketName := terraform.Output(t, terraformOpts, "bucket_name")
	outputAcl := terraform.Output(t, terraformOpts, "acl")
	outputTags := terraform.Output(t, terraformOpts, "tags")

	// asserts
	assert.Contains(t, outputTags, "Environment:dev")
	assert.Contains(t, outputTags, "Managed:terraform")
	assert.Contains(t, outputTags, "Owner:sre")
	assert.Contains(t, outputTags, outputBucketName)
	assert.Equal(t, expectedName, outputBucketName)
	assert.Equal(t, expectedAcl, outputAcl)
}
