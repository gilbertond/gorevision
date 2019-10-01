//variable "access_key" {}
//variable "secret_key" {}
//variable "region" {
//  default = "ap-southeast-2"
//}

//Set provider and location/region
provider "aws" {
  //aws as provider.

  alias = "us-east-2"
  region = "us-east-2"
  //region to deploy to (data-centers). us-east-2 is one in Ohio

  version = "2.21"
  //Neccessary to avoid auto download of new version which may cause breaking changes

//  access_key = var.access_key         Alternative to setting these in bash_profile
}

//Set resource to create. Server in this case
resource "aws_instance" "gil_instance_with_sdk_db" {
  //resource "<PROVIDER>_<TYPE>" "<NAME>"

  provider = "aws"
  //this is a map to a provider on line 2
  ami = "ami-0c55b159cbfafe1f0"
  instance_type = "t2.micro"
  tags = {
    "Name" = "server-with-sdks-and-dbs"
  }

  //  wire in the security group. The security group dependency in this case forces terraform to create the secure_group resource before this resource
  vpc_security_group_ids = [
    aws_security_group.terraform_security_group_instance.id]
  user_data = <<-EOF
              #!/bin/bash
              echo "Hello, World" > index.html
              nohup busybox httpd -f -p 8080 &
              EOF
}

resource "aws_security_group" "terraform_security_group_instance" {
  name = "gil-security-group"

  //  By default, AWS does not allow any incoming or outgoing traffic from an EC2 Instance
  //  Security group to allow the EC2 Instance to receive traffic on port 8080
  ingress {
    from_port = 8080
    to_port = 8080
    protocol = "tcp"
    cidr_blocks = [
      "0.0.0.0/0"]
  }
}

resource "aws_sqs_queue" "terraform_queue" {
  name = "gil_sqs_queue"
  delay_seconds = 90
  max_message_size = 2048
  message_retention_seconds = 86400
  receive_wait_time_seconds = 10

  //  *******Used for a FIFO queue********    [https://www.terraform.io/docs/providers/aws/r/sqs_queue.html]
  //  fifo_queue = true
  //  content_based_deduplication = true

  //  ******Server side encrymption*******
  //  kms_master_key_id = "alias/aws/sqs"
  //  kms_data_key_reuse_period_seconds = 300

  redrive_policy = "{\"deadLetterTargetArn\":\"${aws_sqs_queue.terraform_queue_deadletter.arn}\",\"maxReceiveCount\":4}"
  tags = {
    Environment = "staging"
  }
}

resource "aws_sqs_queue" "terraform_queue_deadletter" {
  name = "gil_DLQ"
  delay_seconds = 90
  max_message_size = 2048
  message_retention_seconds = 86400
  receive_wait_time_seconds = 10
}

resource "aws_sns_topic" "terraform_updates" {
  name = "gil-updates-topic"
}

resource "aws_sns_topic_subscription" "terraform_updates_sqs_target" {
  topic_arn = aws_sns_topic.terraform_updates.arn
  protocol  = "sqs"
  endpoint  = aws_sqs_queue.terraform_queue.arn
}                       