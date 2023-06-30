variable "aws_region" {
  default = "eu-central-1"
}

variable "artifacts_bucket" {
  default = "thesortex-artifacts"
}

variable "budget_amount" {
  default = "5.0"
}

variable "website_image_tag" {
  # Update this to the tag of the image in aws ecr
  default = "7b9322da366583bd1f24d623492576eed2c4436f954e2dbbb7e2ef7a4822903c"
}
