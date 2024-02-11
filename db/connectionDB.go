package db

import (
	"context"
	"fmt"

	"SocialNetworkBackend/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN is the connection object to the database
var MongoCN *mongo.Client
var DatabaseName string

// ConnectDB is the function to connect to the database
func ConnectDB(ctx context.Context) error {
	user 	 := ctx.Value(models.Key("user")).(string)
	password := ctx.Value(models.Key("password")).(string)
	host 	 := ctx.Value(models.Key("host")).(string)
	connStr  := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", user, password, host)

	var clientOptions *options.ClientOptions
	client, err := mongo.Connect(ctx, clientOptions.ApplyURI(connStr))
	if err != nil {
		fmt.Println("Error connecting to the database: " + err.Error())
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("Error pinging the database: " + err.Error())
		return err
	}

	fmt.Println("Connected to the database")
	MongoCN = client
	DatabaseName = ctx.Value(models.Key("database")).(string)
	return nil

}

// CheckConnection is the function to check the connection to the database
func CheckConnection() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	return err == nil
}





