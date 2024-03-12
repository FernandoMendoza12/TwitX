package db

import (
	"context"
	"fmt"

	"github.com/FernandoMendoza12/TwitX/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCn *mongo.Client
var DatabaseName string

func ConnectBD(ctx context.Context) error {
	user := ctx.Value(models.Key("user")).(string)
	password := ctx.Value(models.Key("password")).(string)
	host := ctx.Value(models.Key("host")).(string)
	connectionString := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", user, password, host)

	var clienteOptions = options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(ctx, clienteOptions)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("> Connected to the database")
	MongoCn = client
	DatabaseName = ctx.Value(models.Key("databse")).(string)
	return nil
}

func DatabaseConnected() bool {
	err := MongoCn.Ping(context.TODO(), nil)
	return err == nil
}
