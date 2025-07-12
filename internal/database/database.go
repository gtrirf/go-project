package database

import (
	"fmt"
	"github.com/gtrirf/go-project/config"
	"github.com/gtrirf/go-project/internal/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(dbconfig *config.DBConfig) *gorm.DB {
	sqlinfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbconfig.DBHost, dbconfig.DBPort, dbconfig.DBUsername, dbconfig.DBPassword, dbconfig.DBName,
	)

	db, err := gorm.Open(postgres.Open(sqlinfo), &gorm.Config{})
	service.ErrorPanic(err)
	// fmt.Println("Successfully connected to the database")
	return db
}