//Set provider and location/region
provider "aws" {
  //aws as provider.

  region = "us-west-1"
  //region to deploy to (data-centers). us-east-2 is one in Ohio

  //version = "2.21"                    //Neccessary to avoid auto download of new version which may cause breaking changes

}

//Set resource to create. Server in this case
resource "aws_instance" "gil_instance" {
  //resource "<PROVIDER>_<TYPE>" "<NAME>"

  ami = "ami-0fcdcdb074d2bac5f"
  instance_type = "t2.micro"
}