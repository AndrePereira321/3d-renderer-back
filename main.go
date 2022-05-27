package main

import (
	"log"
	"math/rand"
	"server/database"
	"server/logger"
	"server/server"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func startDatabase() *mongo.Client {
	client, err := database.GetClient()
	if err != nil {
		logger.DisplayMessage("CRITICAL ERROR", "Error connecting to the database: "+err.Error())
		log.Fatal(err)
	}

	return client
}

func main() {
	logger.DisplayMessage("DEBUG", "Initing Server")

	dbClient := startDatabase()
	server := server.NewServer()
	rand.Seed(time.Now().UTC().UnixNano())

	logger.Debug("Server Inited! Starting Listening")

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	defer dbClient.Disconnect(*database.GetClientContext())
}
