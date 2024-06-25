
resource "aws_lambda_function" "golambda" {
  function_name    = "golambda"
  handler          = "golambdabin"
  runtime          = "provided.al2023"
  role             = aws_iam_role.lambda_execution_role.arn
  filename         = "../build/bootstrap.zip"
  source_code_hash = "${base64sha256(filebase64("../build/bootstrap.zip"))}"
  memory_size      = 128
  timeout          = 10
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
