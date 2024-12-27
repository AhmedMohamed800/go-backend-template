package main

import (
	"fmt"
	"log"

	"github.com/AhmedMohamed800/go-backend-template/cmd/registration"
	"github.com/AhmedMohamed800/go-backend-template/config"
	"github.com/AhmedMohamed800/go-backend-template/internal/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config, err := config.LoadConfig()

	if err != nil{
		log.Fatalf("Error loading config: %v", err) // Exit with a fatal error if loading the config fails
	}

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: config.APIConfig.CORSOrigins,  
        AllowMethods: config.APIConfig.CORSMethods,  
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	})) 
	
	db, err := db.NewStorage(config.DBConfig) 
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err) // Exit if database connection fails
	}

	defer db.Close()

	registration.RegisterRoutes(e, db)

    if err := e.Start(fmt.Sprintf(":%s", config.APIConfig.ServerPort)); err != nil {
        log.Fatalf("Error starting server: %v", err) // Exit if the server fails to start
    }
}
