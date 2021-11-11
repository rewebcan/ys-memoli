package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/rewebcan/ys-memoli/internal/usecase/repo"
	"github.com/swaggo/echo-swagger"
)

// NewRouter
// @title YS Case Study
// @description
// @version     1.0
// @BasePath    /
func NewRouter(handler *echo.Echo, settingsRepository repo.SettingsRepository) {
	handler.GET("/swagger/*", echoSwagger.WrapHandler)
	{
		newV1settingsRoutes(handler, settingsRepository)
	}
}
