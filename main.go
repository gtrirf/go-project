package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gtrirf/go-project/config"
	"github.com/gtrirf/go-project/internal/database"
)

func main(){
	router := fiber.New()
	app := fiber.New()
	
	app.Mount("/api", router)

	router.Get("/servercheck", func(c *fiber.Ctx) error {
		return  c.Status(200).JSON(fiber.Map{
			"status":"success",
			"message":"Welcom to Golang, fiber, and Gorm",
		})
	})


	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("❌ Config xato: %v", err)
	}

	db := database.Connect(&cfg)
	database.RunMigrations(db)

	database.RunMigrations(db)
	log.Fatal(app.Listen(":3000"))
}