package handler

import (
	"net/http"

	"github.com/gowok/gowok/web"
	"github.com/ikhsanfalakh/geo-id/internal/service"
)

type LocationHandler struct {
	Service *service.LocationService
}

func NewLocationHandler(s *service.LocationService) *LocationHandler {
	return &LocationHandler{Service: s}
}

func (h *LocationHandler) GetStates(ctx *web.Ctx) error {
	states, err := h.Service.GetStates()
	if err != nil {
		ctx.Res().Status(http.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
		return nil
	}
	ctx.Res().JSON(states)
	return nil
}

func (h *LocationHandler) GetState(ctx *web.Ctx) error {
	id := ctx.Req().Params("id")
	println("GetState called with id:", id)
	state, err := h.Service.GetState(id)
	if err != nil {
		println("GetState error:", err.Error())
		ctx.Res().Status(http.StatusNotFound).JSON(map[string]string{"error": err.Error()})
		return nil
	}
	ctx.Res().JSON(state)
	return nil
}

func (h *LocationHandler) GetCities(ctx *web.Ctx) error {
	id := ctx.Req().Params("id")
	cities, err := h.Service.GetCities(id)
	if err != nil {
		ctx.Res().Status(http.StatusNotFound).JSON(map[string]string{"error": err.Error()})
		return nil
	}
	ctx.Res().JSON(cities)
	return nil
}

func (h *LocationHandler) GetCity(ctx *web.Ctx) error {
	id := ctx.Req().Params("id")
	city, err := h.Service.GetCity(id)
	if err != nil {
		ctx.Res().Status(http.StatusNotFound).JSON(map[string]string{"error": err.Error()})
		return nil
	}
	ctx.Res().JSON(city)
	return nil
}

func (h *LocationHandler) GetDistricts(ctx *web.Ctx) error {
	id := ctx.Req().Params("id")
	districts, err := h.Service.GetDistricts(id)
	if err != nil {
		ctx.Res().Status(http.StatusNotFound).JSON(map[string]string{"error": err.Error()})
		return nil
	}
	ctx.Res().JSON(districts)
	return nil
}

func (h *LocationHandler) GetDistrict(ctx *web.Ctx) error {
	id := ctx.Req().Params("id")
	district, err := h.Service.GetDistrict(id)
	if err != nil {
		ctx.Res().Status(http.StatusNotFound).JSON(map[string]string{"error": err.Error()})
		return nil
	}
	ctx.Res().JSON(district)
	return nil
}

func (h *LocationHandler) GetVillages(ctx *web.Ctx) error {
	id := ctx.Req().Params("id")
	villages, err := h.Service.GetVillages(id)
	if err != nil {
		ctx.Res().Status(http.StatusNotFound).JSON(map[string]string{"error": err.Error()})
		return nil
	}
	ctx.Res().JSON(villages)
	return nil
}

func (h *LocationHandler) GetVillage(ctx *web.Ctx) error {
	id := ctx.Req().Params("id")
	village, err := h.Service.GetVillage(id)
	if err != nil {
		ctx.Res().Status(http.StatusNotFound).JSON(map[string]string{"error": err.Error()})
		return nil
	}
	ctx.Res().JSON(village)
	return nil
}
