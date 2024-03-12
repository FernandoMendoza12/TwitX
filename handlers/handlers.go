package handlers

import (
	"context"
	"fmt"

	"github.com/FernandoMendoza12/TwitX/models"
	"github.com/aws/aws-lambda-go/events"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) models.RespApi {
	fmt.Println("> Processing " + ctx.Value(models.Key("path")).(string) + "\n > " + ctx.Value(models.Key("method")).(string))

	var response models.RespApi
	response.Status = 400

	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {

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
