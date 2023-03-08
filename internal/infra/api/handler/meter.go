package handler

import (
	"net/http"

	"enerBit-system/internal/app"
	"enerBit-system/internal/domain/dto"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Meter interface {
	RegisterNewMeter(cntx echo.Context) error
	DeleteMeter(cntx echo.Context) error
	GetInactiveServiceMeters(cntx echo.Context) error
	GetLastInstallation(cntx echo.Context) error
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

func (handler *meter) DeleteMeter(cntx echo.Context) error {
	ctx := cntx.Request().Context()
	id := cntx.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "meter id is required")
	}
	ID, err := uuid.Parse(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = handler.app.DeleterMeter(ctx, ID); err != nil {
		return err
	}

	return cntx.JSON(http.StatusOK, "Meter Deleted Successfully")
}

func (handler *meter) GetInactiveServiceMeters(cntx echo.Context) error {
	ctx := cntx.Request().Context()
	clientMeters, err := handler.app.GetInactiveServiceMeters(ctx)
	if err != nil {
		return err
	}

	return cntx.JSON(http.StatusOK, clientMeters)
}

func (handler *meter) GetLastInstallation(cntx echo.Context) error {
	ctx := cntx.Request().Context()

	request := dto.LastInstallation{}
	if err := cntx.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := request.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	lastInstallation, err := handler.app.GetLastInstallation(ctx, request)
	if err != nil {
		return err
	}

	return cntx.JSON(http.StatusOK, lastInstallation)
}
