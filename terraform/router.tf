module "router_lambda_ecr" {
  source = "./modules/ecr"

  target_name = "router_lambda"
}