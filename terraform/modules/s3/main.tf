resource "aws_s3_bucket" "bucket" {
  bucket = var.bucket_name
  tags   = {
    Created_By = "Terraform - do not modify in AWS Management Console"
  }
}

resource "aws_s3_bucket_cors_configuration" "bucket_cors" {
  bucket = var.bucket_name

  cors_rule {
    allowed_methods = ["GET"]
    allowed_origins = ["*"]
  }
}

resource "aws_s3_bucket_policy" "allow_access_from_global" {
  bucket = aws_s3_bucket.bucket.id
  policy = jsonencode({
    Version : "2012-10-17",
    Statement : [
      {
        Sid : "PublicReadGetObject",
        Effect : "Allow",
        Principal : "*",
        Action : ["s3:GetObject"],
        Resource : ["arn:aws:s3:::${var.bucket_name}/*"]
      }
    ]
  })
}