package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthController struct{}

func (h *HealthController) CheckHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, struct {
		Status string `json:"status"`
	}{Status: "ok"})
}
