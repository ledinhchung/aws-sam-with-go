package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name := request.PathParameters["name"]

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello, %v", name),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
