# S3 bucket for artifacts to download
module "website_s3_artifacts" {
  source = "./modules/s3"

  bucket_name = var.artifacts_bucket
}

module "website_lambda_ecr" {
  source = "./modules/ecr"

  target_name = "website_lambda"
}
