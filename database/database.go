package database

import (
	"context"
	"server/globals"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/xerrors"
)

var client *mongo.Client = nil
var clientContext *context.Context = nil

func GetClientContext() *context.Context {
	if clientContext == nil {
		context := context.Background()
		clientContext = &context
	}
	return clientContext
}

func GetClientOptions() *options.ClientOptions {
	clientOptions := options.Client()
	clientOptions.ApplyURI(globals.DATABASE_URL)

	return clientOptions
}

func GetClient() (*mongo.Client, error) {
	if client == nil {
		cli, err := mongo.NewClient(GetClientOptions())

		if err != nil {
			return nil, xerrors.Errorf("Error creating the database client: %w", err)
		}

		err = cli.Connect(*GetClientContext())

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
