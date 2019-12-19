package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "OK",
			StatusCode: 200,
		}, nil
	}

	svc := sqs.New(sess)
	sqsUrl := "https://sqs.us-east-1.amazonaws.com/334351750236/myqueue.fifo"

	result, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:            &sqsUrl,
		WaitTimeSeconds:     aws.Int64(20),
		MaxNumberOfMessages: aws.Int64(10),
	})

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Error %v", err),
			StatusCode: 500,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("OK %v", len(result.Messages)),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
