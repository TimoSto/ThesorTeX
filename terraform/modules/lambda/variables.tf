variable "image_url" {
  type = string
}

variable "function_name" {
  type = string
}

variable "lambda_policies" {
  type    = list(string)
  default = []
}

variable "env" {
  type    = map(string)
  default = {
    DUMMY : "DUMMY" //TODO: Get rid of this
  }
}
