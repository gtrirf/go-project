package database

import (
	"log"

	"github.com/gtrirf/go-project/internal/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		&models.Teacher{},
		&models.Group{},
		&models.Student{},
		&models.Payment{},
		&models.MonthlyFee{},
		&models.StudentFee{},
		&models.Attendance{},
		&models.Location{},
		&models.StudentCode{},
	)

	if err != nil {
		log.Fatalf("❌ Migrations error %v", err)
	}
	log.Println("✅ All tables migrated successfully")
}