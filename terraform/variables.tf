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
  default = "5f410ce6"
}
