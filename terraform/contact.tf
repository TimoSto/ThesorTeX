module "contact_lambda_ecr" {
  source = "./modules/ecr"

  target_name = "contact_lambda"
}

module "contact_lambda" {
  source          = "./modules/lambda"
  image_url       = "${module.contact_lambda_ecr.repository_url}:${var.contact_image_tag}"
  function_name   = "contact_lambda"
  dynamo_tables   = [
    aws_dynamodb_table.feedbacks.arn
  ]
}

module "contact_api_gw" {
  source           = "./modules/apigateway"
  api_name         = "contact_api"
  integration_uri  = module.contact_lambda.invoke_arn
  lambda_func_name = module.contact_lambda.func_name
}

resource "aws_dynamodb_table" "feedbacks" {
  name           = "feedbacks"
  billing_mode   = "PROVISIONED"
  read_capacity  = "10"
  write_capacity = "10"
  attribute {
    name = "id"
    type = "S"
  }
  hash_key = "id"
}
