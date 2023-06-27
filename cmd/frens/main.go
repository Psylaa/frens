// cmd/frens/main.go
package main

import (
	"log"

	"github.com/bwoff11/frens/internal/config"
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/router"
	"github.com/bwoff11/frens/internal/storage"
)

func main() {
	// Read the config
	cfg, err := config.ReadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	// Initialize logger
	logger.Init(cfg.Server.LogLevel)

	// Initialize storage
	if err := storage.InitStorage(cfg); err != nil {
		logger.Log.Fatal().Err(err).Msg("Error initializing storage")
	}

	// Connect to the database
	if err := database.InitDB(cfg); err != nil {
		logger.Log.Fatal().Err(err).Msg("Error connecting to database")
	}

	// Initialize router and start the server
	router.Init(cfg.Server.Port, cfg.Server.JWTSecret, cfg.Server.JWTDuration)
}
