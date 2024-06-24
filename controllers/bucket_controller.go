package controllers

import (
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/s3"
  "github.com/aws/aws-sdk-go/service/s3/s3manager"
  "fmt"
  "os"

  "photo-backup-server/utils"
)


func ListBuckets() {
  svc := utils.GetS3Client()

	result, err := svc.ListBuckets(nil)
	if err != nil {
		utils.ExitErrorf("Unable to list buckets, %v", err)
	}

	fmt.Println("Buckets:")

	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
}

func ListBucketObjects() {
  svc := utils.GetS3Client()

  bucket := "photo-backup-travis-linkey"
  resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(bucket)})
  if err != nil {
    utils.ExitErrorf("Unable to list items in bucket %q, %v", bucket, err)
  }

  for _, item := range resp.Contents {
    fmt.Println("Name:         ", *item.Key)
    fmt.Println("Last modified:", *item.LastModified)
    fmt.Println("Size:         ", *item.Size)
    fmt.Println("Storage class:", *item.StorageClass)
    fmt.Println("")
  }
}

func UploadFileToBucket() {
  bucket := "photo-backup-travis-linkey"

  if len(os.Args) != 2 {
    utils.ExitErrorf("Filename required\nUsage: %s filename", os.Args[0])
  }

  filename := os.Args[1]

  file, err := os.Open(filename)
  if err != nil {
    utils.ExitErrorf("Unable to open file %q, %v",  err)
  }

  defer file.Close()

  sess, err := session.NewSession(&aws.Config{
    Region: aws.String("us-west-2")},
  )

  uploader := s3manager.NewUploader(sess)

  _, err = uploader.Upload(&s3manager.UploadInput{
    Bucket: aws.String(bucket),
    Key: aws.String(filename),
    Body: file,
  })

  if err != nil {
    utils.ExitErrorf("Unable to upload %q to %q, %v", filename, bucket, err)
  }

  fmt.Printf("Sucessfully uploaded %q to %q\n", filename, bucket)
}
