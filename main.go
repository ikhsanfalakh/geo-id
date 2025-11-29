package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gowok/gowok"
	"github.com/gowok/gowok/web"
	"github.com/ikhsanfalakh/geo-id/internal/handler"
	"github.com/ikhsanfalakh/geo-id/internal/service"
)

func main() {
	cwd, _ := os.Getwd()
	dataDir := filepath.Join(cwd, "data")
	
	svc := service.NewLocationService(dataDir)
	h := handler.NewLocationHandler(svc)

	// Register routes
	gowok.Web.HttpServeMux.HandleFunc(http.MethodGet, "/test", gowok.Web.Handler.Handler(func(ctx *web.Ctx) error {
		ctx.Res().JSON(map[string]string{"message": "test works"})
		return nil
	}))
	gowok.Web.HttpServeMux.HandleFunc(http.MethodGet, "/test/{id}", gowok.Web.Handler.Handler(func(ctx *web.Ctx) error {
		id := ctx.Req().Params("id")
		ctx.Res().JSON(map[string]string{"id": id})
		return nil
	}))
	
	gowok.Web.HttpServeMux.HandleFunc(http.MethodGet, "/states", gowok.Web.Handler.Handler(h.GetStates))
	gowok.Web.HttpServeMux.HandleFunc(http.MethodGet, "/states/{id}", gowok.Web.Handler.Handler(h.GetState))
	gowok.Web.HttpServeMux.HandleFunc(http.MethodGet, "/states/{id}/cities", gowok.Web.Handler.Handler(h.GetCities))
	
	gowok.Web.HttpServeMux.HandleFunc(http.MethodGet, "/cities/{id}", gowok.Web.Handler.Handler(h.GetCity))
	gowok.Web.HttpServeMux.HandleFunc(http.MethodGet, "/cities/{id}/districts", gowok.Web.Handler.Handler(h.GetDistricts))
	
	gowok.Web.HttpServeMux.HandleFunc(http.MethodGet, "/districts/{id}", gowok.Web.Handler.Handler(h.GetDistrict))
	gowok.Web.HttpServeMux.HandleFunc(http.MethodGet, "/districts/{id}/villages", gowok.Web.Handler.Handler(h.GetVillages))
	
	gowok.Web.HttpServeMux.HandleFunc(http.MethodGet, "/villages/{id}", gowok.Web.Handler.Handler(h.GetVillage))

	gowok.Run()
}
