package main

import (
	"log"
	"os"

	"github.com/BrunoKrugel/go-webhook/internal/api"
	"github.com/BrunoKrugel/go-webhook/internal/client"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err := client.InitMongo()
	if err != nil {
		log.Fatal("Error connecting to MongoDB")
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/:user", api.Webhook)

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
