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

// @title Frens API
// @version 1.0
// @description ActivityPub social network API
// @termsOfService http://swagger.io/terms/

// @contact.name Frens Repo
// @contact.url http://www.github.com/bwoff11/frens

// @license.name MIT License
// @license.url http://www.github.com/bwoff11/frens/docs/LICENSE.md

// @host localhost:3001
// @BasePath /v1
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
	router := router.New(cfg, db, srv)
	router.Run()
}
