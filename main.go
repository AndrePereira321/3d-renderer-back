package main

import (
	"log"
	"math/rand"
	"server/database"
	"server/database/cache"
	"server/globals"
	"server/logger"
	"server/server"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func StartDatabase() *mongo.Client {
	client, err := database.GetClient()
	if err != nil {
		logger.DisplayMessage("CRITICAL ERROR", "Error connecting to the database: "+err.Error())
		log.Fatal(err)
	}

	return client
}

func main() {
	logger.DisplayMessage("DEBUG", "Initing Database")

	dbClient := StartDatabase()

	logger.Debug("Database Inited! Initing server")

	server := server.NewServer()
	rand.Seed(time.Now().UTC().UnixNano())

	err := cache.Cache.Init()
	if err != nil {
		logger.LogError("Error initing cache: "+err.Error(), "main", globals.ERROR_DATABASE_ERROR)
		log.Fatal(err)
		return
	}

	logger.Debug("Server Inited! Starting Listening")

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return
	}

	defer dbClient.Disconnect(*database.GetClientContext())
}
