# S3 bucket for artifacts to download
module "website_s3_artifacts" {
  source = "./modules/s3"

  bucket_name = var.artifacts_bucket
}

module "website_lambda_ecr" {
  source = "./modules/ecr"

  target_name = "website_lambda"
}

module "website_lambda" {
  source          = "./modules/lambda"
  image_url       = "${module.website_lambda_ecr.repository_url}:${var.website_image_tag}"
  function_name   = "website_lambda"
  s3_buckets      = [
    "arn:aws:s3:::${module.website_s3_artifacts.bucket}",
    "arn:aws:s3:::${module.website_s3_artifacts.bucket}/*",
  ]
}

module "website_api_gw" {
  source           = "./modules/apigateway"
  api_name         = "website-api"
  integration_uri  = module.website_lambda.invoke_arn
  lambda_func_name = module.website_lambda.func_name
}

module "website_domain" {
  source      = "./modules/customdomain"
  domain_name = "thesortex.com"
  subdomains  = []
}
