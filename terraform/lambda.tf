locals {
  lambda_file = "../artifacts/website/lambda.zip"
}

resource "aws_lambda_function" "website_lambda_func" {
  filename         = local.lambda_file
  function_name    = "thesortex-website"
  handler          = "prod"
  source_code_hash = base64sha256(local.lambda_file)
  runtime          = "go1.x"
  role             = aws_iam_role.website_lambda_exec.arn
}

data aws_iam_policy_document lambda_assume_role {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

data aws_iam_policy_document lambda_s3 {
  statement {
    actions = [
      "s3:ListBucket",
      "s3:GetObject"
    ]

    resources = [
      "arn:aws:s3:::${aws_s3_bucket.artifacts.bucket}",
      "arn:aws:s3:::${aws_s3_bucket.artifacts.bucket}/*",
    ]
  }
}

resource aws_iam_policy lambda_s3 {
  name        = "lambda-s3-permissions"
  description = "Contains S3 put permission for lambda"
  policy      = data.aws_iam_policy_document.lambda_s3.json
}

resource aws_iam_role website_lambda_exec {
  name               = "website_lambda_exec"
  assume_role_policy = data.aws_iam_policy_document.lambda_assume_role.json
}

resource aws_iam_role_policy_attachment lambda_s3 {
  role       = aws_iam_role.website_lambda_exec.name
  policy_arn = aws_iam_policy.lambda_s3.arn
}

# Attach role to Managed Policy
variable "iam_policy_arn" {
  description = "IAM Policy to be attached to role"
  type        = list(string)

  default = [
    "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
  ]
}

resource "aws_iam_policy_attachment" "role_attach" {
  name       = "policy-thesortex-website"
  roles      = [aws_iam_role.website_lambda_exec.id]
  count      = length(var.iam_policy_arn)
  policy_arn = element(var.iam_policy_arn, count.index)
}