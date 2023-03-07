package app

import (
	"context"
	"net/http"

	"enerBit-system/internal/domain/dto"
	"enerBit-system/internal/domain/ports/postgres/repo"
	"enerBit-system/internal/infra/adapters/postgres/model"
	"github.com/labstack/echo/v4"
)

type Client interface {
	NewInstallation(ctx context.Context, request dto.NewInstallation) error
	GetInstallationByAddress(ctx context.Context, address string) (model.Client, error)
}

type client struct {
	meterApp Meter
	repo     repo.Repository
}

func NewClientApp(meterApp Meter, meterRepo repo.Repository) Client {
	return &client{meterApp, meterRepo}
}

func (app *client) NewInstallation(ctx context.Context, request dto.NewInstallation) error {
	meterDB, err := app.meterApp.GetMeterByBrandAndSerial(ctx, request.Meter.Brand, request.Meter.Serial)
	if err != nil {
		return err
	}

	if meterDB.ID.ID() == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "this meter not exist")
	}

	if meterDB.InUse {
		return echo.NewHTTPError(http.StatusConflict, "this meter is being used at another property")
	}

	clientDB, err := app.GetInstallationByAddress(ctx, request.Address)
	if err != nil {
		return err
	}

	if clientDB.ID != 0 {
		if clientDB.RetirementDate == nil {
			return echo.NewHTTPError(http.StatusBadRequest, "the property already has a meter installed")
		}
		if clientDB.RetirementDate.After(request.InstallationDate) {
			return echo.NewHTTPError(http.StatusConflict, "a meter cannot be installed on this property as its last meter was removed after the current installation date.")
		}
	}

	installation := model.NewInstallation{
		Meter: model.Meter{
			ID:               meterDB.ID,
			InUse:            meterDB.InUse,
			LastInstallation: request.InstallationDate,
		},
		Client: model.Client{
			Address:          request.Address,
			IsActive:         request.IsActive,
			InstallationDate: request.InstallationDate,
			MeterID:          meterDB.ID,
			RetirementDate:   nil,
		},
	}

	if err = app.repo.NewInstallation(ctx, installation); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (app *client) GetInstallationByAddress(ctx context.Context, address string) (model.Client, error) {
	clientDB, err := app.repo.GetInstallationByAddress(ctx, address)
	if err != nil {
		return model.Client{}, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return clientDB, nil
}
