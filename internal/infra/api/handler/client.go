package handler

import (
	"net/http"

	"enerBit-system/internal/app"
	"enerBit-system/internal/domain/dto"
	"github.com/labstack/echo/v4"
)

type Client interface {
	NewInstallation(cntx echo.Context) error
}

type client struct {
	app app.Client
}

func NewClientHandler(app app.Client) Client {
	return &client{app}
}

func (handler *client) NewInstallation(cntx echo.Context) error {
	ctx := cntx.Request().Context()
	request := dto.NewInstallation{}
	if err := cntx.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := request.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := handler.app.NewInstallation(ctx, request); err != nil {
		return err
	}

	return cntx.JSON(http.StatusOK, request)
}
