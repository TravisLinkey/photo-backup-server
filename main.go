package main

import (
  "photo-backup-server/routes"

  "github.com/gin-gonic/gin"
  "fmt"
)



func main() {
  fmt.Println("-- Server Working --")

  r := gin.Default()

  routes.SetupRoutes(r)

  r.Run(":8080")
}
