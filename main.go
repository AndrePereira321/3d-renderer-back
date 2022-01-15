package main

import (
	"context"
	"log"
	"server/database"
	"server/logger"
)

func main() {
	logger.DisplayMessage("DEBUG", "Starting Server")

	client, err := database.GetClient()
	if err != nil {
		logger.DisplayMessage("CRITICAL ERROR", "Error connecting to the database: "+err.Error())
		log.Fatal(err)
	}

	logger.Debug("Server Started!")

	defer client.Disconnect(context.TODO())
}
