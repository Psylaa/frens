// cmd/frends/main.go
package main

import (
	"log"

	"github.com/bwoff11/frens/internal/config"
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/router"
)

func main() {
	// Read the config
	cfg, err := config.ReadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	// Connect to the database
	if err := database.InitDB(cfg); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Initialize router and start the server
	router.Init(cfg.Server.Port, cfg.Server.JWTSecret, cfg.Server.JWTDuration)
}
