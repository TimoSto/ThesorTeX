resource "aws_lambda_function" "lambda_func" {
  function_name = var.function_name
  image_uri     = var.image_url
  package_type  = "Image"
  role          = aws_iam_role.lambda_exec.arn
  environment {
    variables = var.env
  }
}

resource "aws_cloudwatch_log_group" "lambda_func" {
  name = "/aws/lambda/${var.function_name}"

  retention_in_days = 7
}

resource aws_iam_role lambda_exec {
  name               = "lambda_exec-${var.function_name}"
  assume_role_policy = <<EOF
{
 "Version": "2012-10-17",
 "Statement": [
   {
     "Action": "sts:AssumeRole",
     "Principal": {
       "Service": "lambda.amazonaws.com"
     },
     "Effect": "Allow",
     "Sid": ""
   }
 ]
}
EOF
}

resource "aws_iam_policy" "lambda_basic_role" {

  name        = "lambda_basic_policy-${var.function_name}"
  path        = "/"
  description = "AWS IAM Policy for managing aws lambda role"
  policy      = <<EOF
{
 "Version": "2012-10-17",
 "Statement": [
   {
     "Action": [
       "logs:CreateLogStream",
       "logs:PutLogEvents"
     ],
     "Resource": "arn:aws:logs:*:*:*",
     "Effect": "Allow"
   },
  {
     "Action": [
       "logs:CreateLogGroup"
     ],
     "Resource": "arn:aws:logs:*:*:*",
     "Effect": "Deny"
   }
 ]
}
EOF
}

resource aws_iam_role_policy_attachment lambda_basic {
  role       = aws_iam_role.lambda_exec.name
  policy_arn = aws_iam_policy.lambda_basic_role.arn
}

resource "aws_iam_role_policy_attachment" "lambda_basic_execution_role" {
  role       = aws_iam_role.lambda_exec.id
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

# s3 permissions

data aws_iam_policy_document lambda_s3 {
  count = length(var.s3_buckets) > 0 ? 1 : 0

  statement {
    actions = [
      "s3:ListBucket",
      "s3:GetObject"
    ]

    resources = var.s3_buckets
  }
}

resource aws_iam_policy lambda_s3 {
  count = length(var.s3_buckets) > 0 ? 1 : 0

  name        = "lambda-s3-permissions-${var.function_name}"
  description = "Contains S3 put permission for lambda"
  policy      = data.aws_iam_policy_document.lambda_s3[0].json
}

resource aws_iam_role_policy_attachment lambda_s3 {
  count = length(var.s3_buckets) > 0 ? 1 : 0

  role       = aws_iam_role.lambda_exec.name
  policy_arn = aws_iam_policy.lambda_s3[0].arn
}

# dynamodb permissions

data aws_iam_policy_document lambda_dynamodb {
  count = length(var.dynamo_tables) > 0 ? 1 : 0

  statement {
    actions = [
      "dynamoDB:PutItem"
    ]

    resources = var.dynamo_tables
  }
}

resource aws_iam_policy lambda_dynamodb {
  count = length(var.dynamo_tables) > 0 ? 1 : 0

  name        = "lambda-s3-permissions-${var.function_name}"
  description = "Contains S3 put permission for lambda"
  policy      = data.aws_iam_policy_document.lambda_dynamodb[0].json
}

resource aws_iam_role_policy_attachment lambda_dynamodb {
  count = length(var.dynamo_tables) > 0 ? 1 : 0

  role       = aws_iam_role.lambda_exec.name
  policy_arn = aws_iam_policy.lambda_dynamodb[0].arn
}
