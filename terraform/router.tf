module "router_lambda_ecr" {
  source = "./modules/ecr"

  target_name = "router_lambda"
}

module "router_lambda" {
  source        = "./modules/lambda"
  image_url     = "${module.router_lambda_ecr.repository_url}:${var.router_image_tag}"
  function_name = "router_lambda"
  env           = {
    ROUTER_APPS = "WEBSITE",
    ROUTER_TARGET_URL__WEBSITE : aws_apigatewayv2_api.website_lambda.api_endpoint,
    ROUTER_ROUTE__WEBSITE : "/",
  }
  lambda_policies = [
    "arn:aws:s3:::${module.website_s3_artifacts.bucket}",
    "arn:aws:s3:::${module.website_s3_artifacts.bucket}/*",
  ]
}

module "router_api_gw" {
  source          = "./modules/apigateway"
  api_name        = "router_api"
  integration_uri = module.router_lambda.invoke_arn
}
