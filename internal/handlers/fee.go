package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gtrirf/go-project/internal/service"
)

type FeeHandler struct {
	Service *service.FeeService
}

func NewFeeHandler(s *service.FeeService) *FeeHandler {
	return &FeeHandler{Service: s}
}

func (h *FeeHandler) GetFees(c *fiber.Ctx) error {
	fees, err := h.Service.GetAllFees()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error":"information not found"})
	}
	return c.JSON(fees)
}