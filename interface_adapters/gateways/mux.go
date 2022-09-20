package gateways

import (
	"undefeated-davout/echo-api-sample/interface_adapters/controllers"

	"github.com/labstack/echo/v4"
)

func NewMux(e *echo.Echo) {
	e.GET("/health", controllers.HealthController)
}
