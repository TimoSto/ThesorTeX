module "router_lambda_ecr" {
  source = "./modules/ecr"

  target_name = "router_lambda"
}

module "router_lambda" {
  source        = "./modules/lambda"
  image_url     = "${module.router_lambda_ecr.repository_url}:${var.router_image_tag}"
  function_name = "router_lambda"
  env           = {
    ROUTER_APPS = "WEBSITE,CONTACT",
    ROUTER_TARGET_URL__WEBSITE : module.website_api_gw.api_endpoint,
    ROUTER_ROUTE__WEBSITE : "/",
    ROUTER_TARGET_URL__CONTACT : module.contact_api_gw.api_endpoint,
    ROUTER_ROUTE__CONTACT : "/contact/",
  }
}

module "router_api_gw" {
  source          = "./modules/apigateway"
  api_name        = "router_api"
  integration_uri = module.router_lambda.invoke_arn
  lambda_func_name = module.router_lambda.func_name
}
