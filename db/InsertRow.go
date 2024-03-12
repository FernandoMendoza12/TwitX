package db

import (
	"context"

	"github.com/FernandoMendoza12/TwitX/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertRow(user models.User) (string, bool, error) {
	ctx := context.TODO()

	db := MongoCn.Database(DatabaseName)
	col := db.Collection("users")

	user.Password, _ = EncryptPassword(user.Password)

	result, err := col.InsertOne(ctx, user)

	if err != nil {
		return "", false, err
	}

	ObjId, _ := result.InsertedID.(primitive.ObjectID)
	return ObjId.String(), true, nil
}
