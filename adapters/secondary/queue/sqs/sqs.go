package queueadaptersqs

import (
	"fmt"
	"strconv"

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

	// ID           string  `json:"id"`
	// Email        string  `json:"email"`
	// Amount       float64 `json:"amount"`
	// Currency     string  `json:"currency"`
	// ProductID    int64   `json:"productid"`
	// PaymentToken string  `json:"paymenttoken"`
	result, err := svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"id": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(order.ID),
			},
			"email": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(order.Email),
			},
			"amount": &sqs.MessageAttributeValue{
				DataType:    aws.String("Number"),
				StringValue: aws.String(fmt.Sprintf("%f", order.Amount)),
			},
			"currency": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(order.Currency),
			},
			"productid": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(strconv.Itoa(order.ProductID)),
			},
		},
		MessageBody: aws.String("Order"),
		QueueUrl:    &qURL,
	})

	if err != nil {
		fmt.Println("Error while sending message to sqs", err)
	} else {
		fmt.Println("Success while sending message to sqs", *result.MessageId)
	}

	return err
}
