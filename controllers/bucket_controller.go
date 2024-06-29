package controllers

import (
  // "github.com/aws/aws-sdk-go/aws"
  // "github.com/aws/aws-sdk-go/aws/session"
  // "github.com/aws/aws-sdk-go/service/s3"
  // "github.com/aws/aws-sdk-go/service/s3/s3manager"
  "github.com/gin-gonic/gin"

  // "fmt"
  "net/http"

  // "photo-backup-server/utils"
)


func ListBuckets(c *gin.Context) {
  c.IndentedJSON(http.StatusOK, "ListBucket route is working!")
  
  // svc := utils.GetS3Client()

	// result, err := svc.ListBuckets(nil)
	// if err != nil {
	// 	utils.ExitErrorf("Unable to list buckets, %v", err)
	// }

  // c.IndentedJSON(http.StatusOK, result.Buckets)
}

func ListBucketObjects(c *gin.Context) {
  c.IndentedJSON(http.StatusOK, "ListBucketObjects route is working!")

  // svc := utils.GetS3Client()

  // bucket := "photo-backup-travis-linkey"
  // resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(bucket)})
  // if err != nil {
  //   utils.ExitErrorf("Unable to list items in bucket %q, %v", bucket, err)
  // }

  // c.IndentedJSON(http.StatusOK, resp.Contents)
}

func UploadFileToBucket(c *gin.Context) {
  c.IndentedJSON(http.StatusOK, "UploadFileToBucket route is working!")

  // bucket := "photo-backup-travis-linkey"

  // fileHeader, err := c.FormFile("file")
  // if (err != nil) {
  //   utils.ExitErrorf("File not passed in form data")
  // }

  // f, err := fileHeader.Open()
  // if (err != nil) {
  //   utils.ExitErrorf("Error opening file header, %v", err)
  // }

  // sess, err := session.NewSession(&aws.Config{
  //   Region: aws.String("us-west-2")},
  // )

  // uploader := s3manager.NewUploader(sess)

  // _, err = uploader.Upload(&s3manager.UploadInput{
  //   Bucket: aws.String(bucket),
  //   Key: aws.String(fileHeader.Filename),
  //   Body: f,
  // })

  // if err != nil {
  //   utils.ExitErrorf("Unable to upload %q to %q, %v", fileHeader.Filename, bucket, err)
  // }

  // fmt.Printf("Sucessfully uploaded %q to %q\n", fileHeader.Filename, bucket)

  // c.IndentedJSON(http.StatusOK, "Successfully uploaded files!")
}
