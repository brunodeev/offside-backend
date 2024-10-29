package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectToDB() error {
	var err error

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://<user>:<password>@<db-name>.vqn36.mongodb.net/?retryWrites=true&w=majority&appName=<app-name>").SetServerAPIOptions(serverAPI)

	client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		return fmt.Errorf("failed to connect to mongo db: %w", err)
	}

	return nil
}
