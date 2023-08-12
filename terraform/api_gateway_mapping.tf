resource "aws_apigatewayv2_api_mapping" "thesortex" {
  api_id      = module.router_api_gw.id
  domain_name = module.website_domain.id
  stage       = module.router_api_gw.default_stage_id
}

output "custom_domain_api" {
  value = "https://${aws_apigatewayv2_api_mapping.thesortex.domain_name}"
}
