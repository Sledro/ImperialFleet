package main

import (
	"ImperialFleet/api"
	"ImperialFleet/config"
	"ImperialFleet/db"
)

// main - entry point
func main() {
	config.LoadConfig()
	db.ConnectToDatabase()
	api.StartServer()
}
