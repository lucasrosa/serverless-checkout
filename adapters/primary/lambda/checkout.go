package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/lucasrosa/serverless-checkout/businesslogic/cart"
)

// CheckoutAdapter is the interface that defines the entrypoints to this adapter
type CheckoutAdapter interface {
	PlaceOrder(request events.APIGatewayProxyRequest) (Response, error)
}

type checkoutAdapter struct {
	checkoutService cart.CheckoutPrimaryPort
}

func NewCheckoutAdapter(checkoutService cart.CheckoutPrimaryPort) CheckoutAdapter {
	return &checkoutAdapter{
		checkoutService,
	}
}

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

func (a *checkoutAdapter) PlaceOrder(request events.APIGatewayProxyRequest) (Response, error) {

	order := cart.Order{}

	err := json.Unmarshal([]byte(request.Body), &order)

	if err != nil {
		return Response{StatusCode: 400}, nil
	}

	err = a.checkoutService.PlaceOrder(&order)

	if err != nil {
		return Response{StatusCode: 502}, err
	}

	resp := Response{
		StatusCode:      201,
		IsBase64Encoded: false,
		Headers: map[string]string{
			"Content-Type":                     "application/json",
			"Access-Control-Allow-Credentials": "true",
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Methods":     "POST",
			"Access-Control-Allow-Headers":     "application/json",
		},
	}

	return resp, nil
}
