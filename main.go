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

var ginLambda *ginadapter.GinLambdaV2

func init() {
  log.Printf("Gin cold start")
  r := gin.Default()

  routes.SetupRoutes(r)

  ginLambda = ginadapter.NewV2(r)
}


func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
  return ginLambda.ProxyWithContext(ctx, req)
}


func main() {
  lambda.Start(Handler)
}
