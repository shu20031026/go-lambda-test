package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler() {
	fmt.Println("hello go")
}

func main() {
	lambda.Start(handler)
}
