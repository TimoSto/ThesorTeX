resource "aws_ecr_repository" "ecr_repo" {
  name                 = var.target_name
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "aws_ecr_registry_scanning_configuration" "configuration" {
  scan_type = "ENHANCED"
}
