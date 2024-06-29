resource "aws_cloudwatch_log_group" "apigw_logs" {
  name = "/aws/apigateway/photo-backup-server-api"
  retention_in_days = 14
}

resource "aws_iam_role" "apigw_cloudwatch_logs_role" {
  name = "Server-APIGWCloudWatchLogsRole"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole",
        Principal = {
          Service = "apigateway.amazonaws.com"
        },
        Effect = "Allow",
      },
    ]
  })
}

resource "aws_iam_policy" "server_apigw_cloudwatch_logs_policy" {
  name = "Server-APIGWCloudWatchLogsPolicy"
  description = "API Gateway CloudWatch Logs Policy"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:DescribeLogGroups",
          "logs:DescribeLogStreams",
          "logs:PutLogEvents",
          "logs:GetLogEvents",
          "logs:FilterLogEvents",
        ],
        Resource = "*",
        Effect = "Allow",
      },
    ]
  })
}

resource "aws_iam_role_policy_attachment" "apigw_cloudwatch_logs_attachement" {
  policy_arn = aws_iam_policy.server_apigw_cloudwatch_logs_policy.arn
  role = aws_iam_role.apigw_cloudwatch_logs_role.name
}
