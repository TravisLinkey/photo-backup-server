resource "aws_apigatewayv2_api" "photo_backup_api_gateway" {
  name = "photo-backup-api"
  description = "API Gateway V2 for server"
  protocol_type = "HTTP"
  cors_configuration {
    allow_origins = ["*"]
    allow_methods = ["*"]
    allow_headers = ["*"]
    max_age = 300
 }
}

resource "aws_apigatewayv2_route" "route" {
  for_each = { for route in var.routes : route.path => route }

  api_id    = aws_apigatewayv2_api.photo_backup_api_gateway.id
  route_key = "${each.value.path}"
  target    = "integrations/${aws_apigatewayv2_integration.server_integration.id}"
  # authorization_type = each.value.authorization

  # Conditionally set the authorization_id
  # authorizer_id = each.value.authorization == "JWT" ? aws_apigatewayv2_authorizer.server_lambda_authorizer.id : null
}

resource "aws_apigatewayv2_integration" "server_integration" {
  api_id    = aws_apigatewayv2_api.photo_backup_api_gateway.id
  integration_type = "AWS_PROXY"
  integration_method = "POST"
  integration_uri   = aws_lambda_function.photo_backup_server.invoke_arn
  payload_format_version = "2.0"
}

 # Create API Gateway v2 stage (deployment)
resource "aws_apigatewayv2_stage" "server_stage" {
  api_id = aws_apigatewayv2_api.photo_backup_api_gateway.id
  name = "$default"
  description = "prod stage"
  auto_deploy = true

  default_route_settings {
    logging_level = "INFO"
    throttling_rate_limit = 10000
    throttling_burst_limit = 5000
  }

  access_log_settings {
    destination_arn = aws_cloudwatch_log_group.apigw_logs.arn
    format          = jsonencode({
      requestId       = "$context.requestId",
      ip              = "$context.identity.sourceIp",
      caller          = "$context.identity.caller",
      user            = "$context.identity.user",
      requestTime     = "$context.requestTime",
      httpMethod      = "$context.httpMethod",
      resourcePath    = "$context.resourcePath",
      status          = "$context.status",
      protocol        = "$context.protocol",
      responseLength  = "$context.responseLength",
      integrationLatency = "$context.integrationLatency",
      userAgent       = "$context.identity.userAgent",
      requestBody     = "$input.body"
    })
  }
}

# Grant permission to API Gateway to invoke Lambda function
resource "aws_lambda_permission" "resource_name" {
  action = "lambda:InvokeFunction"
  function_name = aws_lambda_function.photo_backup_server.function_name
  principal = "apigateway.amazonaws.com"
  source_arn = "${aws_apigatewayv2_api.photo_backup_api_gateway.execution_arn}/*"
}

