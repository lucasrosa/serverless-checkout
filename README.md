# Serverless Checkout 
[![Build Status](https://travis-ci.org/lucasrosa/serverless-checkout.svg?branch=master)](https://travis-ci.org/lucasrosa/serverless-checkout) 
[![Go Report Card](https://goreportcard.com/badge/github.com/lucasrosa/serverless-checkout)](https://goreportcard.com/report/github.com/lucasrosa/serverless-checkout) 
[![codecov](https://codecov.io/gh/lucasrosa/serverless-checkout/branch/master/graph/badge.svg)](https://codecov.io/gh/lucasrosa/serverless-checkout)


Serverless Checkout is an example of a Serverless Microservice that accepts orders from a client and processes them asynchronously by separating the "checkout" and "process" functions using AWS SQS.

The key features of Serverless Checkout are:
- Automatic scalability
- High availability
- Pay-per-execution
- You can create any stage (dev, test, uat, prod...) you want with a single ```sls deploy -s STAGE``` command
- Infrastructure-as-code, you only have to manage the [serverless.yml](serverless.yml), nothing else! 

## Architecture 

The application uses the following AWS services:
- AWS API Gateway for the REST API
- AWS Lambda for the processing
- AWS SQS to separate the 'checkout' and 'process' functions
- AWS DynamoDB to persist the order
- AWS CloudWatch to save logs

The application is created with the [Serverless Framework](https://serverless.com/) to make it easier to manage and deploy. The entire architecture is described in the [serverless.yml](serverless.yml) file.

#### Architecture diagram
![Alt text](architecture.png?raw=true "Architecture")


## How to run it
### Requirements
- An AWS account
- Go 1.12.x
- Node 10.15.x (required for the Serverless Framework, which is intalled via NPM)
- Serverless Framework >=1.28.0


#### 1. Installing the requirements
- [Go installation guide](https://golang.org/doc/install)
- [NPM/Node](https://nodejs.org/en/)
- [Serverless Framework installation](https://serverless.com/framework/docs/getting-started/)

#### 2. Clone the repository
```git clone https://github.com/lucasrosa/serverless-checkout```

#### 3. Enter into the repository folder
```cd serverless-checkout```

#### 4. Deploying the service
```make deploy```

### Using the service
The output will give you a URL to the REST API of your service. Something like this: https://XXXXXXX.execute-api.us-east-1.amazonaws.com/dev/checkout. 

Now you just have to do a POST to this endpoint with the following parameters as RAW application/json in the body of the request:
```
{
	"id": "some-unique-id",
	"email": "your@email.com",
	"amount": 123.00,
	"currency": "brl",
	"productid": 4,
	"description": "Some product description",
	"paymenttoken": "some-fake-payment-token"
}
```
This POST should return a HTTP 201, indicating that the resource was created with success. After a couple seconds you will be able to see this information pop-up in a table called "orders-STAGE" on your DynamoDB.


## Code
### Business Logic
#### Entities
- Order

#### Public
- Checkout
    - PlaceOrder


### Running tests
```
// In the root folder of the repository run
go test ./...

// To run tests showing the test coverage
go test ./... -cover
```

### Building and deploying the AWS Serverless adapter

#### How to build
```make build```
#### How to deploy
```make deploy``` or ```sls deploy```
#### Deploying a single function
```sls deploy -f checkout```      

### TODO

- [ ] ...