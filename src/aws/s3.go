package aws

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// S3cli definition
type S3Cli struct {
	BucketName string
	S3Session  *s3.S3
	Region     string
}

func NewS3Cli(bucketName string, s3Session *s3.S3, region string) *S3Cli {
	return &S3Cli{BucketName: bucketName, S3Session: s3Session, Region: region}
}

// Starts a s3 session
func S3Session(region string, s3Endpoint string) *s3.S3 {
	mySession := session.Must(session.NewSession())
	return s3.New(mySession, aws.NewConfig().WithRegion(region).WithEndpoint(s3Endpoint).WithS3ForcePathStyle(true))
}

// s3cli can put a new object in s3, the data comes from PorsScanner and print eTag if everything as expected.
func (s3cli *S3Cli) PutObjectToS3(data string) {
	key, _ := os.Hostname()
	input := &s3.PutObjectInput{
		Body:   aws.ReadSeekCloser(strings.NewReader(data)),
		Bucket: aws.String(s3cli.BucketName),
		Key:    aws.String(key),
	}

	r, err := s3cli.S3Session.PutObject(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(r)
}

// Getter objects from S3 returning the content as string
func (s3cli *S3Cli) GetObjectsFromS3(object string) string {
	getObject := &s3.GetObjectInput{
		Bucket: aws.String(s3cli.BucketName),
		Key:    aws.String(object),
	}

	result, err := s3cli.S3Session.GetObject(getObject)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchKey:
				fmt.Println(s3.ErrCodeNoSuchKey, aerr.Error())
			case s3.ErrCodeInvalidObjectState:
				fmt.Println(s3.ErrCodeInvalidObjectState, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return ""
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(result.Body)
	myFileContentAsString := buf.String()
	return myFileContentAsString
}

// This function will list the object in the bucket, calls GetObjectsFromS3 and will return an array with the information node -> port
func (s3cli *S3Cli) ListObjectFromS3() []string {
	s3Values := make([]string, 0)
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(s3cli.BucketName),
	}

	result, err := s3cli.S3Session.ListObjectsV2(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				fmt.Println(s3.ErrCodeNoSuchBucket, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return nil
	}

	for _, v := range result.Contents {
		a := s3cli.GetObjectsFromS3(*v.Key)
		s3Values = append(s3Values, a)
		s3Values = append(s3Values, "\n")
	}
	return s3Values
}
