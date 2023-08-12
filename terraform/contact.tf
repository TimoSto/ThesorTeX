module "contact_lambda_ecr" {
  source = "./modules/ecr"

  target_name = "contact_lambda"
}

module "contact_lambda" {
  source          = "./modules/lambda"
  image_url       = "${module.contact_lambda_ecr.repository_url}:${var.contact_image_tag}"
  function_name   = "contact_lambda"
  lambda_policies = [
    "arn:aws:s3:::${module.website_s3_artifacts.bucket}",
    "arn:aws:s3:::${module.website_s3_artifacts.bucket}/*",
  ]
}

module "contact_api_gw" {
  source           = "./modules/apigateway"
  api_name         = "contact_api"
  integration_uri  = module.contact_lambda.invoke_arn
  lambda_func_name = module.contact_lambda.func_name
}
