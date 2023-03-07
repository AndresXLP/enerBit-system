package handler

import (
	"net/http"

	"enerBit-system/internal/app"
	"enerBit-system/internal/domain/dto"
	"github.com/labstack/echo/v4"
)

type Meter interface {
	RegisterNewMeter(cntx echo.Context) error
}

type meter struct {
	app app.Meter
}

func NewMeterHandler(app app.Meter) Meter {
	return &meter{
		app,
	}
}

func (handler *meter) RegisterNewMeter(cntx echo.Context) error {
	ctx := cntx.Request().Context()
	request := dto.Meter{}
	if err := cntx.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := request.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := handler.app.RegisterNewMeter(ctx, request); err != nil {
		return err
	}

	return cntx.JSON(http.StatusOK, request)
}
