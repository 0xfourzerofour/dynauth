package main

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func Handle(ctx context.Context, input utils.GraphqlEvent) error {

	return nil
}

func main() {
	lambda.Start(Handle)
}
