``` 
$ terraform plan \
   -var 'access_key='$TERRAFORM_AWS_ACCESS_KEY \
   -var 'secret_key='$TERRAFORM_AWS_SECRET_KEY
```

```
$ terraform apply \
  -var 'access_key='$TERRAFORM_AWS_ACCESS_KEY \
  -var 'secret_key='$TERRAFORM_AWS_SECRET_KEY
```

`$ terraform show`

```
$ terraform destroy \
  -var 'access_key='$TERRAFORM_AWS_ACCESS_KEY \
  -var 'secret_key='$TERRAFORM_AWS_SECRET_KEY
```

```
  $ aws s3api create-bucket --bucket=terraform-gil-bucket --region=us-east-1
  $ aws s3 cp movies.json s3://terraform-gil-bucket/v1.0.0/movies.json
```

```
Rollback versions: from 1.0.1.to 1.0.0
   $ terraform apply -var="app_version=1.0.1"
   $ terraform apply -var="app_version=1.0.0"
```