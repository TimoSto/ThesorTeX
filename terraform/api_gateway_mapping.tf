resource "aws_apigatewayv2_api_mapping" "thesortex" {
  api_id      = aws_apigatewayv2_api.website_lambda.id
  domain_name = module.website_domain.id
  stage       = aws_apigatewayv2_stage.default.id
}

output "custom_domain_api" {
  value = "https://${aws_apigatewayv2_api_mapping.thesortex.domain_name}"
}
