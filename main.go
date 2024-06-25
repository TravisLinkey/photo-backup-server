package main

import (
  "fmt"

  // "github.com/gin-gonic/gin"
  "github.com/aws/aws-lambda-go/lambda"

  // "photo-backup-server/routes"
)

type Request struct {
  Action string `json:"action"`
  Data string `json:"data"`
}

type Response struct {
  StatusCode string `json:"statusCode"`
  StatusMessage string `json:"statusMessage"`
}

func handler(event Request) (Response, error) {
  fmt.Println("%v\n", event)

  return Response{StatusCode: "200", StatusMessage: "Success"}, nil
}


func main() {
  fmt.Println("-- Server Working --")

  lambda.Start(handler)

  // r := gin.Default()
  // routes.SetupRoutes(r)

  // r.Run(":8080")
}
