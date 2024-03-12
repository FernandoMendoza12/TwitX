package db

import (
	"context"

	"github.com/FernandoMendoza12/TwitX/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckExistingUser(email string) (models.User, bool, string) {
	ctx := context.TODO()

	db := MongoCn.Database(DatabaseName)
	col := db.Collection("users")

	condition := bson.M{"email": email}

	var result models.User

	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID

}
