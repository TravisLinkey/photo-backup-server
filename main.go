package main

import (
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/s3"
  "fmt"
  "os"
)


func exitErrorf(msg string, args ...interface{}) {
  fmt.Fprintf(os.Stderr, msg + "\n", args...)
  os.Exit(1)
}

func getS3Client() *s3.S3 {
  // Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)
  if err != nil {
    exitErrorf("Unable to create AWS session, %v", err)
  }

	// Create S3 service client
	svc := s3.New(sess)
  return svc
}

func listBucketObjects() {
  svc := getS3Client()

  bucketName := "photo-backup-travis-linkey"
  resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(bucketName)})
  if err != nil {
    exitErrorf("Unable to list items in bucket %q, %v", bucketName, err)
  }

  for _, item := range resp.Contents {
    fmt.Println("Name:         ", *item.Key)
    fmt.Println("Last modified:", *item.LastModified)
    fmt.Println("Size:         ", *item.Size)
    fmt.Println("Storage class:", *item.StorageClass)
    fmt.Println("")
  }
}


func listBuckets() {
  svc := getS3Client()

	result, err := svc.ListBuckets(nil)
	if err != nil {
		exitErrorf("Unable to list buckets, %v", err)
	}

	fmt.Println("Buckets:")

	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
}

func main() {
  fmt.Println("-- Server Working --")

  // listBuckets()
  listBucketObjects()
}
