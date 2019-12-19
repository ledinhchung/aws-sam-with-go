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
		Region: aws.String("us-west-1")},
	)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Error",
			StatusCode: 500,
		}, nil
	}

	svc := sqs.New(sess)
	sqsUrl := "https://sqs.us-east-1.amazonaws.com/334351750236/myqueue.fifo"
	name := request.PathParameters["name"]

	result, err := svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds:           aws.Int64(10),
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
