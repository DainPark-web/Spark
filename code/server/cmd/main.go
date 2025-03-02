package main

import (
	"spark/config"
	"spark/database"
	"spark/server"
)

func main() {
	// Load Environment Variables
	config.LoadEnv()

	// Initialize Database
	database.InitDB()

	// Start Server
	server.StartServer()
}
