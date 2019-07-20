# Serverless Checkout 
[![Build Status](https://travis-ci.org/lucasrosa/serverless-checkout.svg?branch=master)](https://travis-ci.org/lucasrosa/serverless-checkout) 
[![Go Report Card](https://goreportcard.com/badge/github.com/lucasrosa/serverless-checkout)](https://goreportcard.com/report/github.com/lucasrosa/serverless-checkout) 
[![codecov](https://codecov.io/gh/lucasrosa/serverless-checkout/branch/master/graph/badge.svg)](https://codecov.io/gh/lucasrosa/serverless-checkout)

## Architecture

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