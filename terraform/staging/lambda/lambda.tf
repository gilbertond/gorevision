provider "aws" {
  region = "us-east-1"
  alias = "us-east-1"
}

resource "aws_lambda_function" "terraform_lambda_1" {
  function_name = "function1"
  handler = "main"
  role = "${aws_iam_role.run_terraform_ambda.arn}"
  runtime = "go1.x"
  s3_bucket = "terraform-gil-bucket"
  s3_key = "v${var.app_version}/app.zip"
//  s3_key = "v1.0.0/app.zip"

//  vpc_config {
//    security_group_ids = []
//    subnet_ids = []
//  }
}

//IAM Role to determine which access a service [e.g. lambda function] has on AWS services [e.g S3]
//Role has no access policy since its not needed
resource "aws_iam_role" "run_terraform_ambda" {
  name = "terraform_exec_role"
  assume_role_policy = <<EOF
{
"Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_api_gateway_resource" "proxy" {
  parent_id = aws_api_gateway_rest_api.terraform_apigw_app.root_resource_id
  path_part = "{proxy+}"
  //activates proxy behaviour, i.e. this resource will match any request path to the app
  rest_api_id = aws_api_gateway_rest_api.terraform_apigw_app.id
}

resource "aws_api_gateway_method" "proxy_root" {
  rest_api_id = aws_api_gateway_rest_api.terraform_apigw_app.id
  resource_id = aws_api_gateway_rest_api.terraform_apigw_app.root_resource_id
  authorization = "NONE"
  http_method = "ANY"
}

//route http GET requests to the created lambda function
resource "aws_api_gateway_integration" "lambda_root" {
  rest_api_id = aws_api_gateway_rest_api.terraform_apigw_app.id
  resource_id = aws_api_gateway_method.proxy_root.resource_id
  http_method = aws_api_gateway_method.proxy_root.http_method

  //  connection_id = "VPC_LINK" //used if dealing with vpc links
  //  connection_type = "VPC_LINK"  //used if dealing with vpc links. VPC_LINK for private connections between API Gateway and a network load balancer in a VPC. INTERNET by default
  integration_http_method = "GET"
  uri = aws_lambda_function.terraform_lambda_1.invoke_arn
  type = "AWS_PROXY"
  //integration type, in this case it causes api gateway to call into another aws service. i.e. lambda function
}

resource "aws_api_gateway_method" "proxy_get" {
  authorization = "NONE"
  http_method = "ANY"
  resource_id = aws_api_gateway_resource.proxy.id
  rest_api_id = aws_api_gateway_rest_api.terraform_apigw_app.id
}

//route http GET requests to the created lambda function
resource "aws_api_gateway_integration" "lambda" {
  http_method = aws_api_gateway_method.proxy_get.http_method
  resource_id = aws_api_gateway_method.proxy_get.resource_id
  rest_api_id = aws_api_gateway_rest_api.terraform_apigw_app.id

  //  connection_id = "VPC_LINK" //used if dealing with vpc links
  //  connection_type = "VPC_LINK"  //used if dealing with vpc links. VPC_LINK for private connections between API Gateway and a network load balancer in a VPC. INTERNET by default
  integration_http_method = "POST"
  uri = aws_lambda_function.terraform_lambda_1.invoke_arn
  type = "AWS_PROXY"
  //integration type, in this case it causes api gateway to call into another aws service. i.e. lambda function
}

resource "aws_api_gateway_deployment" "deploy" {
  depends_on = [
    aws_api_gateway_integration.lambda_root,
    //    aws_api_gateway_integration.lambda
  ]
  rest_api_id = aws_api_gateway_rest_api.terraform_apigw_app.id
  stage_name = "staging"
}

//Allowing API Gateway to Access Lambda
resource "aws_lambda_permission" "apigw" {
  statement_id = "AllowAPIGatewayInvoke"
  action = "lambda:InvokeFunction"
  function_name = aws_lambda_function.terraform_lambda_1.function_name
  principal = "apigateway.amazonaws.com"

  # The "/*/*" portion grants access from any method on any resource
  # within the API Gateway REST API.
  source_arn = "${aws_api_gateway_rest_api.terraform_apigw_app.execution_arn}/*/*"
}

//version app
variable "app_version" {
}