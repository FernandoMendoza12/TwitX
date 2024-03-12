package routers

import (
	"context"
	"encoding/json"

	"fmt"

	"github.com/FernandoMendoza12/TwitX/db"
	"github.com/FernandoMendoza12/TwitX/models"
)

func Register(ctx context.Context) models.RespApi {
	var user models.User
	var response models.RespApi
	response.Status = 400
	fmt.Println("Register Started")

	body := ctx.Value(models.Key("body")).(string)

	err := json.Unmarshal([]byte(body), &response)
	if err != nil {
		response.Message = err.Error()
		fmt.Println(response.Message)
		return response
	}
	if len(user.Email) == 0 {
		response.Message = "Email must be especified"
		fmt.Println(response.Message)
		return response
	}
	if len(user.Password) < 6 {
		response.Message = "Password must have more than 6 characters"
		fmt.Println(response.Message)
		return response
	}

	_, find, _ := db.CheckExistingUser(user.Email)
	if find {
		response.Message = "There is already existing user with this email"
		fmt.Println(response.Message)
		return response
	}

	_, status, err := db.InsertRow(user)
	if err != nil {
		response.Message = "An error occurred during user registration"
		fmt.Println(response.Message)
		return response
	}

	if !status {
		response.Message = "Error during the insert of the user"
		fmt.Println(response.Message)
		return response
	}
	response.Status = 200
	response.Message = "Register Succesfully"
	fmt.Println(response.Message)
	return response
}
