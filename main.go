package main

import (
	"context"
	"strings"

	"os"

	"github.com/FernandoMendoza12/TwitX/awsgo"
	"github.com/FernandoMendoza12/TwitX/db"
	"github.com/FernandoMendoza12/TwitX/handlers"
	"github.com/FernandoMendoza12/TwitX/models"
	"github.com/FernandoMendoza12/TwitX/secretmanager"
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
		return res, nil
	}

	SecretModel, err := secretmanager.GetSecret(os.Getenv("SecretName"))
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error while taking secretname" + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	path := strings.Replace(request.PathParameters["twittergo"], os.Getenv("UrlPrefix"), "", -1)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), path)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), SecretModel.Username)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), SecretModel.Password)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("host"), SecretModel.Host)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("databse"), SecretModel.Database)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("JWTSign"), SecretModel.JWTSign)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), request.Body)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), os.Getenv("BucketName"))

	//Check connection to mongoDB
	err = db.ConnectBD(awsgo.Ctx)
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error connecting to the database" + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	// Create a Api Response

	apiResponse := handlers.Handler(awsgo.Ctx, request)

	if apiResponse.CustomResp == nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: apiResponse.Status,
			Body:       string(apiResponse.Message),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	} else {
		return apiResponse.CustomResp, nil
	}

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
