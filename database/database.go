package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectToDB(dbUser, dbPassword string) error {
	var err error

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@offside-db.vqn36.mongodb.net/?retryWrites=true&w=majority&appName=offside-db", dbUser, dbPassword)).SetServerAPIOptions(serverAPI)

	Client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		return fmt.Errorf("failed to connect to mongo db: %w", err)
	}

	return nil
}
