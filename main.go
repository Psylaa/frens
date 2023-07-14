// main.go
package main

import (
	"log"

	"github.com/bwoff11/frens/internal/config"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/router"
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
	configuration, err := config.ReadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	// Initialize logger
	logger.Init(configuration.Server.LogLevel)
	configuration.Print() // Print the config. Necessary after logger is initialized

	// Initialize router and start the server
	// This will create a service, which will in turn create a database connection
	router := router.New(configuration)
	router.Run()
}
