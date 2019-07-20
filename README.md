# Serverless Checkout 
[![Build Status](https://travis-ci.org/lucasrosa/serverless-checkout.svg?branch=master)](https://travis-ci.org/lucasrosa/serverless-checkout) 
[![Go Report Card](https://goreportcard.com/badge/github.com/lucasrosa/serverless-checkout)](https://goreportcard.com/report/github.com/lucasrosa/serverless-checkout) 
[![codecov](https://codecov.io/gh/lucasrosa/serverless-checkout/branch/master/graph/badge.svg)](https://codecov.io/gh/lucasrosa/serverless-checkout)


Serverless Checkout is an example of a Serverless Microservice that accepts orders from a client and processes them assynchronously by separating the "checkout" and "process" functions using AWS SQS.

## Architecture 

The service uses the following AWS services:
- AWS API Gateway for the REST API
- AWS Lambda for the processing
- AWS SQS to separate the 'checkout' and 'process' functions
- AWS DynamoDB to persist the order
- AWS CloudWatch to save logs

The entire architecture is described in the [serverless.yml](serverless.yml) file.

#### Architecture diagram
![Alt text](architecture.png?raw=true "Architecture")

## Code
### Business Logic
#### Entities
- Order

#### Public
- Checkout
    - PlaceOrder

### Requirements
#### For the core business logic
- Go 1.12.x
- Remember to set ```export GO111MODULE=on;```

#### For the AWS Serverless adapter
- Serverless Framework >=1.28.0
- An AWS account

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