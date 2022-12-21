resource "aws_s3_bucket" "artifacts" {
  bucket = var.artifacts_bucket
  tags = {
    Created_By = "Terraform - do not modify in AWS Management Console"
  }
}

resource "aws_s3_bucket_cors_configuration" "artifacts_cors" {
  bucket = var.artifacts_bucket

  cors_rule {
    allowed_methods = ["GET"]
    allowed_origins = ["*"]
  }
}

resource "aws_s3_bucket_policy" "allow_access_from_global" {
  bucket = aws_s3_bucket.artifacts.id
  policy = jsonencode({
    Version: "2012-10-17",
    Statement: [{
      Sid: "PublicReadGetObject",
      Effect:"Allow",
      Principal: "*",
      Action:["s3:GetObject"],
      Resource:["arn:aws:s3:::${var.artifacts_bucket}/*"]
    }]
  })
}
