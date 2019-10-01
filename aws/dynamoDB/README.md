DynamoDB: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithItems.html
RUN while setting region `AWS_REGION=us-east-2 go run init-db.go
`
Deploy delete function
`aws lambda create-function --function-name DeleteMovie \
      --zip-file fileb://./deployment.zip \
      --runtime go1.x --handler main \
      --role arn:aws:iam::ACCOUNT_ID:role/DeleteMovieRole \
      --environment Variables={TABLE_NAME=movies} \
      --region us-east-1`