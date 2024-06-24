package utils

import (
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/s3"
)

func GetS3Client() *s3.S3 {
  // Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)
  if err != nil {
    ExitErrorf("Unable to create AWS session, %v", err)
  }

	// Create S3 service client
	svc := s3.New(sess)
  return svc
}


