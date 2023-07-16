// main.go
package main

import (
	"github.com/bwoff11/frens/api/router"
	"github.com/bwoff11/frens/pkg/config"
	"github.com/bwoff11/frens/pkg/database"
	"github.com/bwoff11/frens/service"
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
	config, err := config.Load()
	if err != nil {
		panic(err)
	}

	db, err := database.New(&config.Database)
	if err != nil {
		panic(err)
	}

	service := service.New(db)

	router := router.New(service, &config.API)
	router.Start()
}
