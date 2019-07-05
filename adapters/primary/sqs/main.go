package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	databaseadapterdynamodb "github.com/lucasrosa/serverless-checkout/adapters/secondary/database"
	"github.com/lucasrosa/serverless-checkout/businesslogic/checkout"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, sqsEvent events.SQSEvent) (Response, error) {
	for _, message := range sqsEvent.Records {
		fmt.Printf("The message %s for event source %s = %s \n", message.MessageId, message.EventSource, message.Body)
		//jsonBody, err := json.Marshal(message.Body)
		// if err != nil {
		// 	fmt.Println("err while parsing", err)
		// 	//return Response{StatusCode: 500}, err
		// }

		//fmt.Println("json", string(jsonBody))

		order := checkout.Order{}
		err := json.Unmarshal([]byte(string(message.Body)), &order)

		if err != nil {
			fmt.Println("xerr while converting json to struct", err)
			//return Response{StatusCode: 500}, err
		}
		fmt.Println("order", order)
		//fmt.Println("message", message)

		processRepo := databaseadapterdynamodb.NewDynamoCheckoutRepository()
		processService := checkout.NewProcessService(processRepo)
		err = processService.ProcessOrder(&order)

		if err != nil {
			fmt.Println("error saving in repo", err)
		}
	}

	var buf bytes.Buffer

	fmt.Println("Go Process v1.0! Your function executed successfully!")
	body, err := json.Marshal(map[string]interface{}{
		"message": "Go Process v1.0! Your function executed successfully!",
	})
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "hello-handler",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
