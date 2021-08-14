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

type S3Cli struct {
	BucketName string
	S3Session  *s3.S3
	Region     string
}

func NewS3Cli(bucketName string, s3Session *s3.S3, region string) *S3Cli {
	return &S3Cli{BucketName: bucketName, S3Session: s3Session, Region: region}
}

func S3Session(region string, s3Endpoint string) *s3.S3 {
	mySession := session.Must(session.NewSession(
	/*&aws.Config{
		//Credentials: credentials.NewStaticCredentials("foo", "var", ""),
		Endpoint: aws.String(s3Endpoint),
	}*/))
	return s3.New(mySession, aws.NewConfig().WithRegion(region).WithEndpoint(s3Endpoint).WithS3ForcePathStyle(true))
	//return s3.New(mySession, aws.NewConfig())
}

func (s3cli *S3Cli) PutObjectToS3(data string) {
	//bucketName := os.Getenv("S3_BUCKET")

	key, _ := os.Hostname()
	input := &s3.PutObjectInput{
		Body:   aws.ReadSeekCloser(strings.NewReader(data)),
		Bucket: aws.String(s3cli.BucketName),
		Key:    aws.String(key),
		//ServerSideEncryption: aws.String("AES256"),
		Tagging: aws.String("key1=value1&key2=value2"),
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
	fmt.Println(myFileContentAsString)
	return myFileContentAsString
}

func (s3cli *S3Cli) ListObjectFromS3() []string {
	s3Values := make([]string, 0)
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(s3cli.BucketName),
		//MaxKeys: aws.Int64(2),
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
