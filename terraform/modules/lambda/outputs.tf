output "invoke_arn" {
  value = aws_lambda_function.lambda_func.invoke_arn
}

output "func_name" {
  value = aws_lambda_function.lambda_func.function_name
}
