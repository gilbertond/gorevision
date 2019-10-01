package main

import "github.com/aws/aws-lambda-go/lambda"

type Response struct {
	StatusCode int    `json:"statusCode"`
	Body       string `json:"body"`
}

//A handler func can only have 0-2 args and shud return an error
func handler() (Response, error) {
	return Response{
		StatusCode: 200,
		Body:       "Welcome to Serverless world",
	}, nil
}

func main() {
	//register entrypoint handler which gets executed when the lambda is invoked
	lambda.Start(handler)
}
