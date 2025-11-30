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

// GetStates godoc
// @Summary Get all states
// @Description Get list of all provinces in Indonesia
// @Tags states
// @Produce json
// @Success 200 {array} model.Region
// @Failure 500 {object} model.ErrorResponse
// @Router /states [get]
func (h *LocationHandler) GetStates(c *fiber.Ctx) error {
	states, err := h.Service.GetStates()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(states)
}

// GetState godoc
// @Summary Get state by ID
// @Description Get specific province details by its code
// @Tags states
// @Produce json
// @Param id path string true "State Code (e.g. 11)"
// @Success 200 {object} model.Region
// @Failure 404 {object} model.ErrorResponse
// @Router /states/{id} [get]
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

// GetCities godoc
// @Summary Get cities in state
// @Description Get list of cities/regencies in a specific province
// @Tags states
// @Produce json
// @Param id path string true "State Code (e.g. 11)"
// @Success 200 {array} model.Region
// @Failure 404 {object} model.ErrorResponse
// @Router /states/{id}/cities [get]
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

// GetCity godoc
// @Summary Get city by ID
// @Description Get specific city/regency details by its code
// @Tags cities
// @Produce json
// @Param id path string true "City Code (e.g. 11.01)"
// @Success 200 {object} model.Region
// @Failure 404 {object} model.ErrorResponse
// @Router /cities/{id} [get]
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

// GetDistricts godoc
// @Summary Get districts in city
// @Description Get list of districts (Kecamatan) in a specific city
// @Tags cities
// @Produce json
// @Param id path string true "City Code (e.g. 11.01)"
// @Success 200 {array} model.Region
// @Failure 404 {object} model.ErrorResponse
// @Router /cities/{id}/districts [get]
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

// GetDistrict godoc
// @Summary Get district by ID
// @Description Get specific district details by its code
// @Tags districts
// @Produce json
// @Param id path string true "District Code (e.g. 11.01.01)"
// @Success 200 {object} model.Region
// @Failure 404 {object} model.ErrorResponse
// @Router /districts/{id} [get]
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

// GetVillages godoc
// @Summary Get villages in district
// @Description Get list of villages (Kelurahan/Desa) in a specific district
// @Tags districts
// @Produce json
// @Param id path string true "District Code (e.g. 11.01.01)"
// @Success 200 {array} model.Region
// @Failure 404 {object} model.ErrorResponse
// @Router /districts/{id}/villages [get]
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

// GetVillage godoc
// @Summary Get village by ID
// @Description Get specific village details by its code
// @Tags villages
// @Produce json
// @Param id path string true "Village Code (e.g. 11.01.01.2001)"
// @Success 200 {object} model.Region
// @Failure 404 {object} model.ErrorResponse
// @Router /villages/{id} [get]
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
