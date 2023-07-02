// main.go
package main

import (
	"log"

	"github.com/bwoff11/frens/internal/config"
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/response"
	"github.com/bwoff11/frens/internal/router"
	"github.com/bwoff11/frens/internal/service"
)

func main() {
	// Read the config
	cfg, err := config.ReadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	// Initialize logger
	logger.Init(cfg.Server.LogLevel)

	// Initialize response package
	response.Init(cfg.Server.BaseURL)

	// Connect to the database
	db, err := database.New(cfg)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Create service
	srv := service.New(db, cfg)

	// Initialize router and start the server
	// Router only uses the db temporarily until migration to service is complete
	router := router.NewRouter(cfg, db, srv)
	router.Run()
}
