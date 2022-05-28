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
)

func startDatabase() {
	_, err := database.Init()
	if err != nil {
		logger.DisplayMessage("CRITICAL ERROR", "Error connecting to the database: "+err.Error())
		log.Fatal(err)
		panic(err)
	}
}

func startServer() *server.Server {
	server := server.NewServer()
	rand.Seed(time.Now().UTC().UnixNano())

	err := cache.Cache.Init()
	if err != nil {
		logger.LogError("Error initing cache: "+err.Error(), "main", globals.ERROR_DATABASE_ERROR)
		log.Fatal(err)
		panic(err)
	}

	return server

}

func main() {
	logger.DisplayMessage("DEBUG", "Initing Database")
	startDatabase()
	logger.Debug("Database Inited! Initing server")
	server := startServer()
	logger.Debug("Server Inited! Starting Listening")

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return
	}

	defer func() {
		if database.Client != nil && database.ClientContext != nil {
			database.Client.Disconnect(*database.ClientContext)
		}
	}()
}
