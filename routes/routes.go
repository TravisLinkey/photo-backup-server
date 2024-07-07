package routes

import (
 "github.com/gin-gonic/gin"
  "photo-backup-server/controllers"
)

func SetupRoutes(r *gin.Engine) {
  r.MaxMultipartMemory = 8 << 20
  homeRoutes(r)
  bucketRoutes(r)
}

func homeRoutes(r *gin.Engine) {
  r.GET("/", controllers.Home)
}

func bucketRoutes(r *gin.Engine) {
  bucketGroup := r.Group("/buckets") 
  {
    bucketGroup.GET("/all", controllers.ListBuckets)
    bucketGroup.GET("/directories", controllers.ListDirectories)
    bucketGroup.GET("/objects/all", controllers.ListBucketObjects)
    bucketGroup.GET("/thumbnails/:directory", controllers.GetThumbnails)
    bucketGroup.GET("/presigned-url", controllers.CreatePreSignedURL)
  }
}
