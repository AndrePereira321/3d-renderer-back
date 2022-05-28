package database

import (
	"context"
	"server/globals"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/xerrors"
)

var Client *mongo.Client = nil
var ClientContext *context.Context = nil
var Database *mongo.Database = nil

func GetClientOptions() *options.ClientOptions {
	clientOptions := options.Client()
	clientOptions.ApplyURI(globals.DATABASE_URL)

	return clientOptions
}

func Init() (*mongo.Database, error) {
	context := context.Background()

	cli, err := mongo.NewClient(GetClientOptions())
	if err != nil {
		return nil, xerrors.Errorf("Error creating the database client: %w", err)
	}

	err = cli.Connect(context)
	if err != nil {
		return nil, xerrors.Errorf("Error connecting to the database: %w", err)
	}

	Client = cli
	ClientContext = &context
	Database = cli.Database(globals.DATABASE_NAME)

	return Database, err
}
