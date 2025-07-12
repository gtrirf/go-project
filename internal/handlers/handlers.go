package handlers

import (
	"fmt"

	// "github.com/gofiber/fiber/v2"
	m "github.com/gtrirf/go-project/internal/models"
	"gorm.io/gorm"
)
func TestGorm(db *gorm.DB) {
		monthly := m.MonthlyFee{
		Amount:250000,
		Description: "Regular",
	}
	db.Create(&monthly)
	fmt.Println("amount successfully created")	
}

// func GetMonthlyFees(c *fiber.Ctx, db *gorm.DB) (err error){
// 	var fees []m.MonthlyFee
// 	if result := db.Find(&fees); result.Error != nil {
// 		return c.Status(404).JSON(fiber.Map{
// 			"error":"Ma'lumotlar topilmadi",
// 		})
// 	}
// 	return c.JSON((fees))
// }