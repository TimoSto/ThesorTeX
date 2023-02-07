locals {
  lambda_file = "../artifacts/website/lambda.zip"
}

resource "aws_lambda_function" "website_lambda_func" {
  filename         = local.lambda_file
  function_name    = "thesortex-website"
  handler          = "cmd"
  source_code_hash = base64sha256(local.lambda_file)
  runtime          = "go1.x"
  role             = aws_iam_role.website_lambda_exec.arn
}

# Assume role setup
resource "aws_iam_role" "website_lambda_exec" {
  name_prefix = "thesortex-website"

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