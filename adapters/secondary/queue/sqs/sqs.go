package queueadaptersqs

import (
	"fmt"

	"github.com/lucasrosa/serverless-checkout/businesslogic/checkout"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type checkoutRepository struct{}

func NewSQSCheckoutRepository() checkout.SecondaryPort {
	return &checkoutRepository{}
}

func (r *checkoutRepository) SendOrderForProcessing(order *checkout.Order) error {
	fmt.Println(order)

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	svc := sqs.New(sess)

	// URL to our queue
	qURL := "https://sqs.us-east-1.amazonaws.com/587998505259/PaymentQueue"

	result, err := svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"Title": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String("The Whistler"),
			},
			"Author": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String("John Grisham"),
			},
			"WeeksOn": &sqs.MessageAttributeValue{
				DataType:    aws.String("Number"),
				StringValue: aws.String("6"),
			},
		},
		MessageBody: aws.String("Information about current NY Times fiction bestseller for week of 12/11/2016."),
		QueueUrl:    &qURL,
	})

	if err != nil {
		fmt.Println("Error while sending message to sqs", err)
	} else {
		fmt.Println("Success while sending message to sqs", *result.MessageId)
	}

	return err
}
