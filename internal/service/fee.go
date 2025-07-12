package service

import (
	"github.com/gtrirf/go-project/internal/models"
	"gorm.io/gorm"
)

type FeeService struct {
	DB *gorm.DB
}

func NewFeeService(db *gorm.DB) * FeeService {
	return &FeeService{DB: db}
}

func (s *FeeService) GetAllFees() ([]models.MonthlyFee, error) {
	var fees []models.MonthlyFee
	result := s.DB.Find(&fees)
	return fees, result.Error
}