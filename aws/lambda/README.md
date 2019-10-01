bundle into zip for upload into AWS lambda

```
    $ GOOS=linux GOARCH=amd64 go build -o main main.go
    $ zip helloworld.zip helloworld
```

Deploy zip to lambda
`aws lambda update-function-code --function-name HelloServerless \
     --zip-file fileb://./deployment.zip \
     --region us-east-1`
     
 Curl
 `curl -sX POST -d '{"id":6, "name": "Spiderman:Homecoming"}' https://51cxzthvma.execute-api.us-east-1.amazonaws.com/staging/movies | jq '.'`

 