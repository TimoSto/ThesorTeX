resource "aws_apigatewayv2_api" "api" {
  name          = var.api_name
  protocol_type = "HTTP"
  description   = "Serverlessland API Gwy HTTP API and AWS Lambda function"

  cors_configuration {
    allow_credentials = false
    allow_headers     = []
    allow_methods     = [
      "GET",
      "HEAD",
      "OPTIONS",
      "POST",
    ]
    allow_origins = [
      "*",
    ]
    expose_headers = []
    max_age        = 0
  }
}

resource "aws_apigatewayv2_domain_name" "domain_name" {
  domain_name = var.domain

  domain_name_configuration {
    certificate_arn = var.cert_arn
    endpoint_type   = "REGIONAL"
    security_policy = "TLS_1_2"
  }
}

resource "aws_cloudwatch_log_group" "api_gw" {
  name = "/aws/api_gw/${aws_apigatewayv2_api.api.name}"

  retention_in_days = 7
}

resource "aws_apigatewayv2_stage" "default" {
  api_id = aws_apigatewayv2_api.api.id

  name        = "$default"
  auto_deploy = true

  access_log_settings {
    destination_arn = aws_cloudwatch_log_group.api_gw.arn

    format = jsonencode({
      requestId               = "$context.requestId"
      sourceIp                = "$context.identity.sourceIp"
      requestTime             = "$context.requestTime"
      protocol                = "$context.protocol"
      httpMethod              = "$context.httpMethod"
      resourcePath            = "$context.resourcePath"
      routeKey                = "$context.routeKey"
      status                  = "$context.status"
      responseLength          = "$context.responseLength"
      integrationErrorMessage = "$context.integrationErrorMessage"
    }
    )
  }
  depends_on = [aws_cloudwatch_log_group.api_gw]
}

resource "aws_apigatewayv2_api_mapping" "api" {
  api_id      = aws_apigatewayv2_api.api.id
  domain_name = var.domain_name_id
  stage       = aws_apigatewayv2_stage.default.id
}

resource "aws_apigatewayv2_integration" "api" {
  for_each = var.integrations

  api_id = aws_apigatewayv2_api.api.id

  integration_uri  = each.integration_uri
  integration_type = "AWS_PROXY"
}

resource "aws_apigatewayv2_route" "api" {
  for_each = var.integrations

  api_id    = aws_apigatewayv2_api.api.id
  route_key = each.key
  target    = "integrations/${aws_apigatewayv2_integration.api[each.key].id}"
}