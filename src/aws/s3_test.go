package aws

import (
	"os"
	"testing"
)

const S3Endpoint = "http://127.0.0.1:4566"
const AWSRegion = "us-east-1"
const BucketName = "test"

func TestS3(t *testing.T) {
	key, _ := os.Hostname()
	session := S3Session(AWSRegion, S3Endpoint)
	s3cli := NewS3Cli(BucketName, session, AWSRegion)
	s3cli.PutObjectToS3("expectedOutput")
	output := s3cli.GetObjectsFromS3(key)
	if output == "expectedOutput" {
		t.Logf("Success !")
	} else {
		t.Errorf("Failed ! got %v want expectedOutput", output)
	}

}
