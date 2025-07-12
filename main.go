package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gtrirf/go-project/config"
	"github.com/gtrirf/go-project/internal/database"
	"github.com/gtrirf/go-project/internal/middleware"
	"github.com/gtrirf/go-project/internal/routers"
	"go.uber.org/zap"
	// "github.com/gtrirf/go-project/internal/handlers"
	// "github.com/gtrirf/go-project/internal/models"
)

func main() {
	app := fiber.New()
	api := app.Group("/api")

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	app.Use(middleware.LoggerMiddleware(logger))


	// connent with .env file
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("❌ Config xato: %v", err)
	}

	// database connections
	db := database.Connect(&cfg)

	// go run check
	api.Get("/servercheck", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcom to Golang, fiber, and Gorm",
		})
	})

	// api.Get("/fees", func(c *fiber.Ctx) error {
	// 	var fees []models.MonthlyFee
	// 	if result := db.Find(&fees); result.Error != nil {
	// 		return c.Status(500).JSON(fiber.Map{"error": "Ma'lumot yo'q"})
	// 	}
	// 	return c.Status(200).JSON(fees)
	// })

	// test gorm
	// handlers.TestGorm(db)

	// database migrations
	// database.RunMigrations(db)

	// api endpoints
	routers.SetupRoutes(app, db)

	log.Fatal(app.Listen(":3000"))
}
