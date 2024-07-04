package controllers

import (
  "log"
  "fmt"
  "net/http"
  "time"

  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/s3"
  "github.com/aws/aws-sdk-go/service/s3/s3manager"
  "github.com/gin-gonic/gin"

  "photo-backup-server/utils"
)

var bucket string = "photo-backup-travis-linkey"

func CreatePreSignedURL(c *gin.Context) {
  subFolder := c.Request.Header.Get("X-Foldername")
  filename := c.Request.Header.Get("X-Filename")
  svc := utils.GetS3Client()

  req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
    Bucket: aws.String(bucket),
    Key: aws.String(subFolder + "/" + filename),
  })

  urlStr, _ := req.Presign(15 * time.Minute)
  c.IndentedJSON(http.StatusOK, gin.H{"url": urlStr})
}

func ListBuckets(c *gin.Context) {
  svc := utils.GetS3Client()

	result, err := svc.ListBuckets(nil)
	if err != nil {
		utils.ExitErrorf("Unable to list buckets, %v", err)
	}

  c.IndentedJSON(http.StatusOK, result.Buckets)
}

func ListBucketObjects(c *gin.Context) {
  svc := utils.GetS3Client()

  resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(bucket)})
  if err != nil {
    utils.ExitErrorf("Unable to list items in bucket %q, %v", bucket, err)
  }

  c.IndentedJSON(http.StatusOK, resp.Contents)
}

func UploadFileToBucket(c *gin.Context) {
  log.Printf("Uploading to bucket")

  fileHeader, err := c.FormFile("file")
  if (err != nil) {
    utils.ExitErrorf("File not passed in form data")
  }

  f, err := fileHeader.Open()
  if (err != nil) {
    utils.ExitErrorf("Error opening file header, %v", err)
  }

  sess, err := session.NewSession(&aws.Config{
    Region: aws.String("us-west-2")},
  )
  if (err != nil) {
    utils.ExitErrorf("Error creating session, %v", err)
  }

  uploader := s3manager.NewUploader(sess)

  _, err = uploader.Upload(&s3manager.UploadInput{
    Bucket: aws.String(bucket),
    Key: aws.String(fileHeader.Filename),
    Body: f,
  })

  if err != nil {
    utils.ExitErrorf("Unable to upload %q to %q, %v", fileHeader.Filename, bucket, err)
  }

  fmt.Printf("Sucessfully uploaded %q to %q\n", fileHeader.Filename, bucket)

  c.IndentedJSON(http.StatusOK, "Successfully uploaded files!")
}
