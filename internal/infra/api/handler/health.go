package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Health struct {
	Code    int    `json:"status"`
	Message string `json:"message"`
}

// HealthCheck godoc
//
//	@Tags			Health
//	@Summary		Check if service is active
//	@Description	health service
//	@Produce		json
//	@Success		200	{object}	Health
//	@Router			/health [get]
func HealthCheck(context echo.Context) error {
	response := &Health{
		Code:    http.StatusOK,
		Message: "Active!",
	}

	return context.JSON(http.StatusOK, response)
}
