variable "api_name" {
  type = string
}

variable "domain" {
  type = string
}

variable "domain_name_id" {
  type = string
}

variable "cert_arn" {
  type = string
}

variable "integrations" {
  type = map(any)
}
