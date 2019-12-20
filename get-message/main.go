package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess)
	sqsUrl := os.Getenv("SQS_URL")

	result, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:            &sqsUrl,
		WaitTimeSeconds:     aws.Int64(0),
		MaxNumberOfMessages: aws.Int64(10),
	})

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Error %v", sqsUrl),
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
