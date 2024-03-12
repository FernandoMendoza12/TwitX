package handlers

import (
	"context"
	"fmt"

	"github.com/FernandoMendoza12/TwitX/jwt"
	"github.com/FernandoMendoza12/TwitX/models"
	"github.com/FernandoMendoza12/TwitX/routers"
	"github.com/aws/aws-lambda-go/events"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) models.RespApi {
	fmt.Println("> Processing " + ctx.Value(models.Key("path")).(string) + "\n > " + ctx.Value(models.Key("method")).(string))

	var response models.RespApi
	response.Status = 400

	isOk, statusCode, msg, _ := validateAuthorization(ctx, request)

	if !isOk {
		response.Status = statusCode
		response.Message = msg
		return response
	}

	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {
		case "register":
			return routers.Register(ctx)
		}
		//
	case "GET":
		switch ctx.Value(models.Key("path")).(string) {

		}
		//
	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {

		}
		//
	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {

		}
		//
	}
	response.Message = "Invalid Method"

	return response
}

func validateAuthorization(ctx context.Context, request events.APIGatewayProxyRequest) (bool, int, string, models.Claim) {
	path := ctx.Value(models.Key("path")).(string)

	if path == "register" || path == "login" || path == "getAvatar" || path == "getBanner" {
		return true, 200, "", models.Claim{}
	}

	token := request.Headers["Authorization"]

	if len(token) == 0 {
		return false, 401, "Token Required", models.Claim{}
	}

	claim, isOk, msg, err := jwt.ProcessToken(token, ctx.Value(models.Key("jwtsign")).(string))
	if !isOk {
		if err != nil {
			fmt.Println("Token error" + err.Error())
			return false, 401, err.Error(), models.Claim{}
		} else {
			fmt.Println("Token error" + msg)
			return false, 401, msg, models.Claim{}
		}
	}

	fmt.Println("Valid Token")
	return true, 200, msg, *claim
}
