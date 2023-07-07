resource "aws_lambda_function" "lambda_func" {
  function_name = var.function_name
  image_uri     = var.image_url
  package_type  = "Image"
  role          = aws_iam_role.lambda_exec.arn
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

resource aws_iam_role lambda_exec {
  name               = "lambda_exec-${var.function_name}"
  assume_role_policy = data.aws_iam_policy_document.lambda_assume_role.json
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
  name       = "policy-${var.function_name}"
  roles      = [aws_iam_role.lambda_exec.id]
  count      = length(var.iam_policy_arn)
  policy_arn = element(var.iam_policy_arn, count.index)
}

data aws_iam_policy_document lambda_s3 {
  statement {
    actions = [
      "s3:ListBucket",
      "s3:GetObject"
    ]

    resources = var.lambda_policies
  }
}

resource aws_iam_policy lambda_s3 {
  name        = "lambda-s3-permissions-${var.function_name}"
  description = "Contains S3 put permission for lambda"
  policy      = data.aws_iam_policy_document.lambda_s3.json
}

resource aws_iam_role_policy_attachment lambda_s3 {
  role       = aws_iam_role.lambda_exec.name
  policy_arn = aws_iam_policy.lambda_s3.arn
}
