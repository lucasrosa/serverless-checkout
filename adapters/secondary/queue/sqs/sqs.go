package queueadaptersqs

import (
	"encoding/json"
	"fmt"

	"github.com/lucasrosa/serverless-checkout/businesslogic/cart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type checkoutRepository struct{}

// NewSQSCheckoutRepository instantiates the repository for this adapter
func NewSQSCheckoutRepository() cart.CheckoutSecondaryPort {
	return &checkoutRepository{}
}

func (r *checkoutRepository) SendOrderForProcessing(order *cart.Order) error {
	fmt.Println(order)

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	svc := sqs.New(sess)

	// URL to our queue
	// FIXME: make this dynamic
	qURL := "https://sqs.us-east-1.amazonaws.com/587998505259/PaymentQueue"

	orderJson, err := json.Marshal(order)
	if err != nil {
		return err
	}

	result, err := svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageBody:  aws.String(string(orderJson)),
		QueueUrl:     &qURL,
	})

	if err != nil {
		fmt.Println("Error while sending message to sqs", err)
	} else {
		fmt.Println("Success while sending message to sqs", *result.MessageId)
	}

	return err
}
