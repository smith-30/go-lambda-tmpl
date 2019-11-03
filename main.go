package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context) error {
	fmt.Printf("%#v\n", "hello")
	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
