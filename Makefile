.PHONY: build clean deploy

build:
	env GO111MODULE=on GOOS=linux go build -ldflags="-s -w" -o bin/checkout adapters/primary/lambda/*
	env GO111MODULE=on GOOS=linux go build -ldflags="-s -w" -o bin/process adapters/primary/sqs/*

test: 
	env GO111MODULE=on go test ./... -cover
	
clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
