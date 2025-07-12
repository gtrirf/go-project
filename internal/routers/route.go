package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gtrirf/go-project/internal/handlers"
	"github.com/gtrirf/go-project/internal/service"
	"gorm.io/gorm"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api")
	
	feeService := service.NewFeeService(db)
	feeHandler := handlers.NewFeeHandler(feeService)

	api.Get("/fees", feeHandler.GetFees)
}