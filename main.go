package main

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"

	"github.com/ikhsanfalakh/geo-id/docs"
	"github.com/ikhsanfalakh/geo-id/internal/handler"
	"github.com/ikhsanfalakh/geo-id/internal/model"
	"github.com/ikhsanfalakh/geo-id/internal/service"
)

// @title Geo-ID API
// @version 1.0
// @description API for Indonesian Administrative Regions (Provinces, Cities, Districts, Villages)
// @host localhost:8080
// @BasePath /
// @tag.name states
// @tag.description Operations regarding provinces
// @tag.name cities
// @tag.description Operations regarding cities/regencies
// @tag.name districts
// @tag.description Operations regarding districts
// @tag.name villages
// @tag.description Operations regarding villages
func main() {
	// Load .env file if it exists (ignore error if file doesn't exist)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables or defaults")
	}

	// Get configuration from environment variables with defaults
	appName := getEnv("APP_NAME", "Geo-ID API")
	appVersion := getEnv("APP_VERSION", "1.0")
	env := getEnv("ENV", "development")
	enableSwagger := getEnvAsBool("ENABLE_SWAGGER", true)

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		AppName: appName + " v" + appVersion,
	})

	// Get data directory
	dataDir := getEnv("DATA_DIR", "")
	if dataDir == "" {
		cwd, _ := os.Getwd()
		dataDir = filepath.Join(cwd, "data")
	}

	// Initialize service and handler
	svc := service.NewLocationService(dataDir)
	h := handler.NewLocationHandler(svc)

	// Configure Swagger host dynamically based on PORT
	port := getEnv("PORT", "8080")
	docs.SwaggerInfo.Host = "localhost:" + port

	// Serve static assets
	app.Static("/assets", "./docs/assets")

	// Swagger route (conditionally enabled)
	if enableSwagger {
		app.Get("/apidocs/*", swagger.New(swagger.Config{
			CustomStyle: `
			.swagger-ui .topbar { background-color: #1b1b1b; }
			.swagger-ui .topbar .link img { display: none; }
			.swagger-ui .topbar .link svg { display: none; }
			.swagger-ui .topbar .link::before { 
				content: '';
				display: inline-block;
				width: 35px;
				height: 35px;
				background-image: url('/assets/logo-geo.svg');
				background-size: contain;
				background-repeat: no-repeat;
				vertical-align: middle;
				margin-right: 10px;
			}
			.swagger-ui .topbar .link::after { 
				content: 'Geo-ID'; 
				color: #fff; 
				font-size: 24px; 
				font-weight: 800; 
				letter-spacing: 1px;
				vertical-align: middle;
			}
			.swagger-ui .topbar .download-url-wrapper { display: none; } 
		`,
		}))
		log.Printf("Swagger UI enabled at /apidocs/index.html")
	} else {
		log.Printf("Swagger UI is disabled (ENV=%s)", env)
	}

	// Register routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(model.NewSuccessResponse(fiber.Map{
			"app":      appName,
			"version":  appVersion,
			"docs_url": "/apidocs/index.html",
			"message":  "Welcome to Geo-ID API",
		}))
	})

	app.Get("/states", h.GetStates)
	app.Get("/states/:id", h.GetState)
	app.Get("/states/:id/cities", h.GetCities)

	app.Get("/cities/:id", h.GetCity)
	app.Get("/cities/:id/districts", h.GetDistricts)

	app.Get("/districts/:id", h.GetDistrict)
	app.Get("/districts/:id/villages", h.GetVillages)

	app.Get("/villages/:id", h.GetVillage)

	// Start server
	log.Printf("Starting %s v%s on port %s (ENV=%s)", appName, appVersion, port, env)
	if err := app.Listen(":" + port); err != nil {
		panic(err)
	}
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsBool retrieves an environment variable as a boolean or returns a default value
func getEnvAsBool(key string, defaultValue bool) bool {
	valStr := getEnv(key, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}
	return defaultValue
}
