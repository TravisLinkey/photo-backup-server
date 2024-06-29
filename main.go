package main

import (
  "log"
  "context"

  "github.com/aws/aws-lambda-go/events"
  "github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
  "github.com/gin-gonic/gin"

  "photo-backup-server/routes"
)

type Request struct {
  Action string `json:"action"`
  Data string `json:"data"`
}

type Response struct {
  StatusCode string `json:"statusCode"`
  StatusMessage string `json:"statusMessage"`
}

var ginLambda *ginadapter.GinLambdaV2

func init() {
  log.Printf("Gin cold start")
  r := gin.Default()

  routes.SetupRoutes(r)

  ginLambda = ginadapter.NewV2(r)

  // r.Run(":8080")
}


func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
  return ginLambda.ProxyWithContext(ctx, req)
}


func main() {
  lambda.Start(Handler)
}
