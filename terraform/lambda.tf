locals {

  lambda_policy_attachements = [
    "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole",
    "arn:aws:iam::aws:policy/AWSXrayWriteOnlyAccess",
    "arn:aws:iam::aws:policy/AmazonS3FullAccess",
  ]

  lambda_file      = "../artifacts/website/lambda.zip"
  source_code_hash = filebase64sha256(local.lambda_file)

}

resource "aws_iam_role" "lambda_service_role" {
  name = "thesortex-website-service"
  path = "/thesortex-website-service/"

  # which entity can assume this role
  assume_role_policy = data.aws_iam_policy_document.lambda_service_role.json
}

# cf. siehe https://www.terraform.io/docs/providers/aws/d/iam_policy_document.html
data "aws_iam_policy_document" "lambda_service_role" {
  statement {
    effect = "Allow"

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }

    actions = ["sts:AssumeRole"]
  }
}

resource "aws_lambda_function" "service" {
  filename         = local.lambda_file
  source_code_hash = local.source_code_hash
  function_name    = "thesortex-website-service"
  role             = aws_iam_role.lambda_service_role.arn
  handler          = "lambda"
  runtime          = "go1.x"
  memory_size      = "128"
  publish          = true
  timeout          = 30

  # cf. https://docs.aws.amazon.com/lambda/latest/dg/enabling-x-ray.html
  tracing_config {
    mode = "Active"
  }

  tags = {
    Name       = "thesortex-website-service-lambda-fn"
    Created_By = "Terraform - do not modify in AWS Management Console"
  }
}
