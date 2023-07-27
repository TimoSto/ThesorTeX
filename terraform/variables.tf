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
  default = "8bbe2f0a"
}

variable "contact_image_tag" {
  # Update this to the tag of the image in aws ecr
  default = "c3eb0be5"
}
