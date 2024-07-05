resource "aws_s3_bucket" "s3_bucket" {
  bucket = "${var.bucket_name}"
}

resource "aws_s3_bucket_policy" "bucket_policy" {
  bucket = aws_s3_bucket.s3_bucket.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow",
        Principal = {
          AWS = [
            aws_iam_role.lambda_execution_role.arn,
            var.thumbnail_execution_role_arn
          ]
        },
        Action = [
          "s3:*"
        ],
        Resource = [
          "arn:aws:s3:::${aws_s3_bucket.s3_bucket.id}",
          "arn:aws:s3:::${aws_s3_bucket.s3_bucket.id}/*"
        ]
      },
      {
        Effect: "Deny",
        Principal: "*",
        Action: "s3:PutObject",
        Resource = [
          "arn:aws:s3:::${aws_s3_bucket.s3_bucket.id}",
          "arn:aws:s3:::${aws_s3_bucket.s3_bucket.id}/*"
        ],
        Condition: {
          StringNotEquals: {
            "s3:x-amz-server-side-encryption": "aws:kms"
          }
        }
      },
    ]
  })
}

resource "aws_s3_bucket_notification" "bucket_notification" {
  bucket = aws_s3_bucket.s3_bucket.id

  lambda_function {
    lambda_function_arn = var.thumbnail_server_arn
    events = ["s3:ObjectCreated:*"]
  }
}
