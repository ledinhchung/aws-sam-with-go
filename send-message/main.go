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
	name := request.PathParameters["name"]

	result, err := svc.SendMessage(&sqs.SendMessageInput{
		MessageBody:            aws.String("This is message send from labda"),
		MessageGroupId:         aws.String(name),
		MessageDeduplicationId: aws.String(name),
		QueueUrl:               &sqsUrl,
	})

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Error %v", err),
			StatusCode: 500,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("OK, %v", result.MessageId),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
