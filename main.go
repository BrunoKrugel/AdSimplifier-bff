package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/BrunoKrugel/go-webhook/internal/api"
	"github.com/BrunoKrugel/go-webhook/internal/client"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	if os.Getenv("ENVIRONMENT") != "prod" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
	}

	err := client.InitMongo()
	if err != nil {
		log.Fatal("Error connecting to MongoDB")
	}

	// Echo instance
	app := echo.New()

	// Middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	// Routes
	app.POST("/:user", api.Webhook)
	app.GET("/:user", api.Webhook)

	if os.Getenv("PORT") == "" {
		log.Fatal("$PORT must be set")
	}

	// Start server
	go func() {
		app.Logger.Fatal(app.Start(":" + os.Getenv("PORT")))
	}()

	// Listen for system signals to gracefully stop the application
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	switch <-signalChannel {
	case os.Interrupt:
		client.CloseMongo()
		log.Println("Received SIGINT, stopping...")
	case syscall.SIGTERM:
		client.CloseMongo()
		log.Println("Received SIGTERM, stopping...")
	}
}
