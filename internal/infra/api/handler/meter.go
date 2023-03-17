package handler

import (
	"net/http"

	"enerBit-system/internal/app"
	"enerBit-system/internal/domain/dto"
	"github.com/andresxlp/gosuite/errs"
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

// @Tags			Meter
// @Summary		Register New Meter
// @Description	Register New Meter
// @Produce		json
// @Param			request	body		dto.Meter	true	"Request Body"
// @Success		200		{object}	dto.Meter
// @Failure		400
// @Failure		404
// @Router			/meter [post]
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
		httpErr, ok := err.(*errs.AppError)
		if ok {
			return httpErr.NewEchoHttpError()
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return cntx.JSON(http.StatusOK, request)
}

// @Tags			Meter
// @Summary		Delete a Meter
// @Description	Delete a Meter
// @Produce		json
// @Param			id	path		string	true	"meter_id"
// @Success		200	{object}	string	"Meter Delete Successfully"
// @Failure		400
// @Failure		404
// @Router			/meter/{id} [delete]
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

// @Tags			Meter
// @Summary		Get Inactive service Meter
// @Description	Get Inactive service Meter
// @Produce		json
// @Success		200	{object}	dto.MeterWithoutService
// @Failure		400
// @Failure		404
// @Router			/meter/inactive [get]
func (handler *meter) GetInactiveServiceMeters(cntx echo.Context) error {
	ctx := cntx.Request().Context()
	clientMeters, err := handler.app.GetInactiveServiceMeters(ctx)
	if err != nil {
		return err
	}

	return cntx.JSON(http.StatusOK, clientMeters)
}

// @Tags			Meter
// @Summary		Get Last Installation Meter
// @Description	Get Last Installation Meter
// @Produce		json
// @Param			brand	query		string	true	"brand meter"
// @Param			serial	query		string	true	"serial meter"
// @Success		200		{object}	dto.Client
// @Failure		400
// @Failure		404
// @Router			/meter/last-installation [get]
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
