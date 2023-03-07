package app

import (
	"context"
	"net/http"

	"enerBit-system/internal/domain/dto"
	"enerBit-system/internal/domain/ports/postgres/repo"
	"enerBit-system/internal/infra/adapters/postgres/model"
	"github.com/labstack/echo/v4"
)

type Meter interface {
	RegisterNewMeter(ctx context.Context, request dto.Meter) error
	GetMeterByBrandAndSerial(ctx context.Context, brand, serial string) (model.Meter, error)
}

type meter struct {
	meterRepo repo.Repository
}

func NewMeterApp(meterRepo repo.Repository) Meter {
	return &meter{meterRepo}
}

func (app *meter) RegisterNewMeter(ctx context.Context, request dto.Meter) error {
	reqModel := model.Meter{}
	reqModel.BuildModel(request)

	_, err := app.GetMeterByBrandAndSerial(ctx, request.Brand, request.Serial)
	if err != nil {
		return err
	}

	if err := app.meterRepo.RegisterNewMeter(ctx, reqModel); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (app *meter) GetMeterByBrandAndSerial(ctx context.Context, brand, serial string) (model.Meter, error) {
	existingMeter, err := app.meterRepo.GetMeterByBrandAndSerial(ctx, brand, serial)
	if err != nil {
		return model.Meter{}, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if existingMeter.ID.ID() != 0 {
		return model.Meter{}, echo.NewHTTPError(http.StatusConflict, "this meter brand and serial already exist")
	}

	return existingMeter, nil
}
