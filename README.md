# Serverless Checkout 

##Architecture
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