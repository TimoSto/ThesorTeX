module "contact_lambda_ecr" {
  source = "./modules/ecr"

  target_name = "contact_lambda"
}
