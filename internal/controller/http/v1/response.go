package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

var loggerFunc = func(err error, statusCode int, msg string) {
	log.Error().Err(err).Int("code", statusCode).Str("message", msg).Send()
}

type successResponse struct {
	Data   interface{} `json:"data,omitempty"`
	Status string      `json:"status,omitempty"`
}

type errResponse struct {
	Error string `json:"message"`
}

func responseError(c echo.Context, code int, err error, msg string) error {
	loggerFunc(err, code, msg)
	return c.JSON(code, errResponse{msg})
}

func responseSuccess(c echo.Context, code int, value interface{}) error {
	if value == nil {
		return c.JSON(code, successResponse{value, "success"})
	}
	return c.JSON(code, successResponse{value, ""})
}
