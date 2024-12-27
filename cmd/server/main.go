package main

import (
	"fmt"
	"log"

	"github.com/AhmedMohamed800/go-backend-template/cmd/setup"
	"github.com/AhmedMohamed800/go-backend-template/config"
	"github.com/AhmedMohamed800/go-backend-template/internal/db"
	"github.com/labstack/echo/v4"
)

func main() {
	config, err := config.LoadConfig()

	if err != nil{
		log.Fatalf("Error loading config: %v", err) // Exit with a fatal error if loading the config fails
	}

	e := echo.New()

	setup.SetupMiddleware(e, config)

	db, err := db.NewStorage(config.DBConfig) 
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err) // Exit if database connection fails
	}

	defer db.Close()

	setup.RegisterRoutes(e, db)

    if err := e.Start(fmt.Sprintf(":%s", config.APIConfig.ServerPort)); err != nil {
        log.Fatalf("Error starting server: %v", err) // Exit if the server fails to start
    }
}


