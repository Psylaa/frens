// main.go
package main

import (
	"log"

	"github.com/bwoff11/frens/internal/config"
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/router"
)

func main() {
	// Read the config
	cfg, err := config.ReadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	// Initialize logger
	logger.Init(cfg.Server.LogLevel)

	// Connect to the database
	db, err := database.New(cfg)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Initialize router and start the server
	router := router.NewRouter(cfg, db)
	router.Run()
}
