// cmd/frends/main.go
package main

import (
	"log"

	"github.com/bwoff11/frens/internal/config"
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/router"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func main() {
	// Read the config
	cfg, err := config.ReadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	// Connect to the database
	database.InitDB(cfg)

	// Initialize router and start the server
	router.Init(cfg.Server.Port, cfg.Server.JWTSecret)
}
