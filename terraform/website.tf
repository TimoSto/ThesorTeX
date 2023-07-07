module "website_s3_artifacts" {
  source = "./modules/s3"

  bucket_name = var.artifacts_bucket
}
