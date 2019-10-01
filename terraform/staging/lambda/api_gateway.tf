//container for other api gatway objects to be created
resource aws_api_gateway_rest_api "terraform_apigw_app" {
  name = "app"
  description = "terraform created api gateway to app"
}

output "base_url" {
  value = aws_api_gateway_deployment.deploy.invoke_url
}