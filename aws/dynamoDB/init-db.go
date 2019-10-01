package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbattribute"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Movie struct {
	ID   int
	Name string
}

func main() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(readData(cfg))

	//movies, err := readMovies("movies.json")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for _, movie := range movies {
	//	fmt.Println("Inserting:", movie.Name)
	//	err = insertMovie(cfg, movie)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}
}

func readMovies(fileName string) ([]Movie, error) {
	movies := make([]Movie, 0)
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return movies, err
	}

	err = json.Unmarshal(data, &movies)
	if err != nil {
		return movies, err
	}

	return movies, nil
}

func insertMovie(cfg aws.Config, movie Movie) error {
	item, err := dynamodbattribute.MarshalMap(movie)
	if err != nil {
		return err
	}

	svc := dynamodb.New(cfg)
	req := svc.PutItemRequest(&dynamodb.PutItemInput{
		TableName: aws.String("movie"),
		Item:      item,
	})

	_, er := req.Send(context.TODO())
	if er != nil {
		return er
	}

	return nil
}

func readData(cfg aws.Config) (string, error) {
	svc := dynamodb.New(cfg)
	req := svc.ScanRequest(&dynamodb.ScanInput{
		//TableName: aws.String(os.Getenv("movie")),
		TableName: aws.String("movie"),
	})

	resp, err := req.Send(context.TODO())
	if err == nil {
		return "Nothing...", err
	}

	items, err := json.Marshal(resp.Items)
	fmt.Println("***********\n", resp.Items)
	return string(items), nil
}

// returning response objects more like responseEntity in java
func findAll() (events.APIGatewayProxyResponse, error) {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while retrieving AWS credentials",
		}, nil
	}

	svc := dynamodb.New(cfg)
	req := svc.ScanRequest(&dynamodb.ScanInput{
		TableName: aws.String("movie"),
	})
	res, err := req.Send(context.TODO())
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while scanning DynamoDB",
		}, nil
	}

	//response, err := json.Marshal(res.Items)

	//returing above line will return raw dynamo struct, may wanna read these into a dto
	movies := make([]Movie, 0)
	for _, item := range res.Items {
		i, _ := strconv.Atoi(*item["ID"].S)
		movies = append(movies, Movie{
			ID:   i,
			Name: *item["Name"].S,
		})
	}
	response, err := json.Marshal(movies)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while decoding to string value",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(response),
	}, nil
}

func findOne(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := request.PathParameters["id"]

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while retrieving AWS credentials",
		}, nil
	}

	svc := dynamodb.New(cfg)
	req := svc.GetItemRequest(&dynamodb.GetItemInput{
		TableName: aws.String("movie"),
		Key: map[string]dynamodb.AttributeValue{
			"ID": dynamodb.AttributeValue{
				S: aws.String(id),
			},
		},
	})
	res, err := req.Send(context.TODO())
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while fetching movie from DynamoDB",
		}, nil
	}

	response, err := json.Marshal(res)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while decoding to string value",
		}, nil
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(response),
	}, nil
}

func insert(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var movie Movie

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while retrieving AWS credentials",
		}, nil
	}

	svc := dynamodb.New(cfg)
	req := svc.PutItemRequest(&dynamodb.PutItemInput{
		TableName: aws.String("movie"),
		Item: map[string]dynamodb.AttributeValue{
			"ID": dynamodb.AttributeValue{
				S: aws.String(string(movie.ID)),
			},
			"Name": dynamodb.AttributeValue{
				S: aws.String(movie.Name),
			},
		},
	})
	_, err = req.Send(context.TODO())
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while inserting movie to DynamoDB",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 204,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string("Created"),
	}, nil
}

func delete(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var movie Movie
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while retrieving AWS credentials",
		}, nil
	}

	svc := dynamodb.New(cfg)
	req := svc.DeleteItemRequest(&dynamodb.DeleteItemInput{
		//aws.String(os.Getenv("TABLE_NAME")) handy for a generic delete/insert/etc
		TableName: aws.String("movie"),
		Key: map[string]dynamodb.AttributeValue{
			"ID": dynamodb.AttributeValue{
				S: aws.String(string(movie.ID)),
			},
		},
	})
	_, err = req.Send(context.TODO())
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while deleting movie from DynamoDB",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusNoContent,
		Body:       "Successfully deleted record",
	}, nil
}
