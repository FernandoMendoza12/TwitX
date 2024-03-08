package main

import (
	"context"

	"os"

	"github.com/FernandoMendoza12/TwitX/awsgo"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(execLambda)

}

func execLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var res *events.APIGatewayProxyResponse
	awsgo.InitAWS()
	if ValidateParams() {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error taking enviroments variables from AWS must have 'SecretName', 'BucketName' and UrlPrefix",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
	}
	return res, nil
}

func ValidateParams() bool {
	_, param := os.LookupEnv("SecretName")
	if !param {
		return param
	}

	_, param = os.LookupEnv("BucketName")
	if !param {
		return param
	}

	_, param = os.LookupEnv("UrlPrefix")
	if !param {
		return param
	}

	return param
}
