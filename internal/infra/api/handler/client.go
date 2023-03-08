package handler

import (
	"net/http"

	"enerBit-system/internal/app"
	"enerBit-system/internal/domain/dto"
	"github.com/labstack/echo/v4"
)

type Client interface {
	NewInstallation(cntx echo.Context) error
	UninstallMeter(cntx echo.Context) error
}

type client struct {
	app app.Client
}

func NewClientHandler(app app.Client) Client {
	return &client{app}
}

// @Tags			Installation
// @Summary		Installation meter in property
// @Description	Install meter in property
// @Produce		json
// @Param			request	body		dto.NewInstallation	true	"Request Body"
// @Success		200		{object}	dto.NewInstallation
// @Failure		400
// @Failure		404
// @Router			/client [post]
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

// @Tags			Installation
// @Summary		Uninstall meter in property
// @Description	Uninstall meter in property
// @Produce		json
// @Param			request	body		dto.UninstallMeter	true	"Request Body"
// @Success		200		{object}	string
// @Failure		400
// @Failure		404
// @Router			/client [delete]
func (handler *client) UninstallMeter(cntx echo.Context) error {
	ctx := cntx.Request().Context()

	request := dto.UninstallMeter{}
	if err := cntx.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := request.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := handler.app.UninstallMeter(ctx, request); err != nil {
		return err
	}
	return cntx.JSON(http.StatusOK, "Meter uninstalled successfully")
}
