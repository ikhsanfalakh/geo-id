package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikhsanfalakh/geo-id/internal/service"
)

type LocationHandler struct {
	Service *service.LocationService
}

func NewLocationHandler(s *service.LocationService) *LocationHandler {
	return &LocationHandler{Service: s}
}

func (h *LocationHandler) GetStates(c *fiber.Ctx) error {
	states, err := h.Service.GetStates()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(states)
}

func (h *LocationHandler) GetState(c *fiber.Ctx) error {
	id := c.Params("id")
	state, err := h.Service.GetState(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(state)
}

func (h *LocationHandler) GetCities(c *fiber.Ctx) error {
	id := c.Params("id")
	cities, err := h.Service.GetCities(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(cities)
}

func (h *LocationHandler) GetCity(c *fiber.Ctx) error {
	id := c.Params("id")
	city, err := h.Service.GetCity(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(city)
}

func (h *LocationHandler) GetDistricts(c *fiber.Ctx) error {
	id := c.Params("id")
	districts, err := h.Service.GetDistricts(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(districts)
}

func (h *LocationHandler) GetDistrict(c *fiber.Ctx) error {
	id := c.Params("id")
	district, err := h.Service.GetDistrict(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(district)
}

func (h *LocationHandler) GetVillages(c *fiber.Ctx) error {
	id := c.Params("id")
	villages, err := h.Service.GetVillages(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(villages)
}

func (h *LocationHandler) GetVillage(c *fiber.Ctx) error {
	id := c.Params("id")
	village, err := h.Service.GetVillage(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(village)
}
