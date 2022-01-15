package database

import (
	"context"
	"server/globals"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/xerrors"
)

var client *mongo.Client = nil

func clientContext() context.Context {
	context, _ := context.WithTimeout(context.Background(), 10*time.Second)

	return context
}

func clientOptions() *options.ClientOptions {
	clientOptions := options.Client()
	clientOptions.ApplyURI(globals.DATABASE_URL)

	return clientOptions
}

func GetClient() (*mongo.Client, error) {
	if client == nil {
		cli, err := mongo.NewClient(clientOptions())

		if err != nil {
			return nil, xerrors.Errorf("Error creating the database client: %w", err)
		}

		ctx := clientContext()
		err = cli.Connect(ctx)

		if err != nil {
			return nil, xerrors.Errorf("Error connecting to the database: %w", err)
		}

		client = cli
	}

	return client, nil
}

func GetDatabase() (*mongo.Database, error) {
	client, err := GetClient()
	if err != nil {
		return nil, xerrors.Errorf("Error retrieving the database client: %w", err)
	}
	return client.Database(globals.DATABASE_NAME), nil
}
