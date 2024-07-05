
resource "aws_lambda_function" "photo_backup_server" {
  function_name    = "photo-backup-server"
  description      = "A lambda to handle creating and s3 presigned urls"
  handler          = "bootstrap"
  runtime          = "provided.al2023"
  role             = aws_iam_role.lambda_execution_role.arn
  filename         = "../build/bootstrap.zip"
  source_code_hash = "${base64sha256(filebase64("../build/bootstrap.zip"))}"
  memory_size      = 128
  timeout          = 10
}

resource "aws_lambda_permission" "apigw_permission" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.photo_backup_server.function_name
  principal     = "apigateway.amazonaws.com"

  # The /*/* portion grants access from any deployment stage of the API Gateway
  source_arn = "${aws_apigatewayv2_api.photo_backup_api_gateway.execution_arn}/*/*"
}

# IAM Role for the Lambda function to execute
resource "aws_iam_role" "lambda_execution_role" {
  name = "lambda_execution_role"

  assume_role_policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Action" : "sts:AssumeRole",
        "Effect" : "Allow",
        "Principal" : {
          "Service" : "lambda.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_role_policy" "lambda_s3_permissions" {
  name = "lambda-s3-policy"
  role = aws_iam_role.lambda_execution_role.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid: "AllowS3Access",
        Effect = "Allow",
        Action = [
          "s3:*",
        ],
        Resource = [
          "arn:aws:s3:::*",      # List specific bucket
          "arn:aws:s3:::*/*"     # Allow actions on objects in the specific bucket
        ]
      },
      {
        Action = [
          "s3:ListAllMyBuckets"
        ],
        Effect = "Allow",
        Resource = "*"
      }
    ]
  })
}

# Attach the necessary policies to the IAM role (Example policy, adjust as needed)
resource "aws_iam_role_policy_attachment" "lambda_role_attachment" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
  role       = aws_iam_role.lambda_execution_role.name
}
