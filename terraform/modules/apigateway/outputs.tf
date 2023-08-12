output "id" {
  value = aws_apigatewayv2_api.api.id
}

output "default_stage_id" {
  value = aws_apigatewayv2_stage.default.id
}

output "api_endpoint" {
  value = aws_apigatewayv2_api.api.api_endpoint
}