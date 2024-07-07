package controllers

import (
  "log"
  "net/http"
  "time"
  "os"

  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/service/s3"
  "github.com/gin-gonic/gin"
  "github.com/joho/godotenv"

  "photo-backup-server/utils"
)

var sourceBucket string
var destBucket string
var s3Client *s3.S3

func init() {
  s3Client = utils.GetS3Client()

  err := godotenv.Load(".env")
  if err != nil {
    log.Fatalf("Error loading .env file")
  }

  sourceBucket = os.Getenv("SOURCE_BUCKET")
  destBucket = os.Getenv("DEST_BUCKET")
}

func CreatePreSignedURL(c *gin.Context) {
  subFolder := c.Request.Header.Get("X-Foldername")
  filename := c.Request.Header.Get("X-Filename")

  req, _ := s3Client.PutObjectRequest(&s3.PutObjectInput{
    Bucket: aws.String(sourceBucket),
    Key: aws.String(subFolder + "/" + filename),
    // ServerSideEncryption: aws.String("aws:kms"),
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

func ListDirectories(c *gin.Context) {
  resp, err := s3Client.ListObjectsV2(&s3.ListObjectsV2Input{
    Bucket: aws.String(sourceBucket),
    Delimiter: aws.String("/"),
  })
  if err != nil {
    utils.ExitErrorf("Unable to list items in bucket %q, %v", sourceBucket, err)
  }

  var directories []string
  for _, prefix := range resp.CommonPrefixes {
    directories = append(directories, *prefix.Prefix)
  }

  c.IndentedJSON(http.StatusOK, directories)
}

func ListBucketObjects(c *gin.Context) {
  resp, err := s3Client.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(sourceBucket)})
  if err != nil {
    utils.ExitErrorf("Unable to list items in bucket %q, %v", sourceBucket, err)
  }

  c.IndentedJSON(http.StatusOK, resp.Contents)
}

func GetThumbnails(c *gin.Context) {
  directory := c.Param("directory")
  if directory == "" {
    utils.ExitErrorf("Query string not provided for directory")
  } 

  log.Printf("Getting thumbnails for directory: %s", directory)

  result, err := s3Client.ListObjectsV2(&s3.ListObjectsV2Input{
    Bucket: aws.String(destBucket),
    Prefix: aws.String(directory),
  })
  if err != nil {
    utils.ExitErrorf("Unable to list thumbnails for %q, %v", destBucket, err)
  }

  var thumbnails []string
  for _, item := range result.Contents {
    req, _ := s3Client.GetObjectRequest(&s3.GetObjectInput{
      Bucket: aws.String(destBucket),
      Key: aws.String(*item.Key),
    })
    urlStr, err := req.Presign(15 * time.Minute)
    if err != nil {
      log.Println("Unable to sign request", err)
      continue
    }
    thumbnails = append(thumbnails, urlStr)
  }

  c.IndentedJSON(http.StatusOK, gin.H{"thumbnails": thumbnails})
}
