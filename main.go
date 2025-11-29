package main

import (
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/ikhsanfalakh/geo-id/internal/handler"
	"github.com/ikhsanfalakh/geo-id/internal/service"
)

func main() {
	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		AppName: "Geo-ID API v1.0",
	})

	// Get data directory
	cwd, _ := os.Getwd()
	dataDir := filepath.Join(cwd, "data")

	// Initialize service and handler
	svc := service.NewLocationService(dataDir)
	h := handler.NewLocationHandler(svc)

	// Register routes
	app.Get("/states", h.GetStates)
	app.Get("/states/:id", h.GetState)
	app.Get("/states/:id/cities", h.GetCities)

	app.Get("/cities/:id", h.GetCity)
	app.Get("/cities/:id/districts", h.GetDistricts)

	app.Get("/districts/:id", h.GetDistrict)
	app.Get("/districts/:id/villages", h.GetVillages)

	app.Get("/villages/:id", h.GetVillage)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	println("Server starting on port", port)
	if err := app.Listen(":" + port); err != nil {
		panic(err)
	}
}
