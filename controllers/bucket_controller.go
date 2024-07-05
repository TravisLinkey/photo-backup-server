package controllers

import (
  "log"
  "net/http"
  "time"

  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/service/s3"
  "github.com/gin-gonic/gin"

  "photo-backup-server/utils"
)

var sourceBucket string = "photo-backup-travis-linkey"
var destBucket string = "photo-backup-thumbnails"

var s3Client *s3.S3

func init() {
  s3Client = utils.GetS3Client()
}

func CreatePreSignedURL(c *gin.Context) {
  subFolder := c.Request.Header.Get("X-Foldername")
  filename := c.Request.Header.Get("X-Filename")

  req, _ := s3Client.PutObjectRequest(&s3.PutObjectInput{
    Bucket: aws.String(sourceBucket),
    Key: aws.String(subFolder + "/" + filename),
    ServerSideEncryption: aws.String("aws:kms"),
  })

  urlStr, _ := req.Presign(15 * time.Minute)
  c.IndentedJSON(http.StatusOK, gin.H{"url": urlStr})
}

func ListBuckets(c *gin.Context) {
	result, err := s3Client.ListBuckets(nil)
	if err != nil {
		utils.ExitErrorf("Unable to list buckets, %v", err)
	}

  c.IndentedJSON(http.StatusOK, result.Buckets)
}

func ListBucketObjects(c *gin.Context) {
  resp, err := s3Client.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(sourceBucket)})
  if err != nil {
    utils.ExitErrorf("Unable to list items in bucket %q, %v", sourceBucket, err)
  }

  c.IndentedJSON(http.StatusOK, resp.Contents)
}

func GetThumbnails(c *gin.Context) {
  log.Printf("Getting thumbnails")

  result, err := s3Client.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(destBucket),})
  if err != nil {
    utils.ExitErrorf("Unable to list thumbnails for %q, %v", destBucket, err)
  }

  var thumbnails []string
  for _, item := range result.Contents {
    req, _ := s3Client.GetObjectRequest(&s3.GetObjectInput{
      Bucket: aws.String(destBucket),
      Key: aws.String(*item.Key),
    })
    urlStr, err := req.Presign(15 * 60)
    if err != nil {
      log.Println("Unable to sign request", err)
      continue
    }
    thumbnails = append(thumbnails, urlStr)
  }

  c.IndentedJSON(http.StatusOK, gin.H{"thumbnails": thumbnails})
}
