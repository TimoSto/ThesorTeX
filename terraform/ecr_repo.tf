moved {
  from = aws_ecr_repository.website
  to   = module.website_lambda_ecr.aws_ecr_repository.ecr_repo
}
resource "aws_ecr_repository" "contact" {
  name                 = "contact_lambda"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}
