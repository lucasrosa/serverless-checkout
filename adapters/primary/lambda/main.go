package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	queueadaptersqs "github.com/lucasrosa/serverless-checkout/adapters/secondary/queue/sqs"
	"github.com/lucasrosa/serverless-checkout/businesslogic/checkout"
)

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request events.APIGatewayProxyRequest) (Response, error) {
	checkoutRepo := queueadaptersqs.NewSQSCheckoutRepository()
	checkoutService := checkout.NewCheckoutService(checkoutRepo)
	checkoutAdapter := NewCheckoutAdapter(checkoutService)

	return checkoutAdapter.PlaceOrder(request)
}

func main() {
	lambda.Start(Handler)
}
