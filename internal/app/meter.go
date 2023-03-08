package app

import (
	"context"
	"net/http"

	"enerBit-system/internal/domain/dto"
	"enerBit-system/internal/domain/ports/postgres/repo"
	"enerBit-system/internal/infra/adapters/postgres/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Meter interface {
	RegisterNewMeter(ctx context.Context, request dto.Meter) error
	GetMeterByBrandAndSerial(ctx context.Context, brand, serial string) (model.Meter, error)
	GetMeterByID(ctx context.Context, ID uuid.UUID) (model.Meter, error)
	DeleterMeter(ctx context.Context, ID uuid.UUID) error
	GetInactiveServiceMeters(ctx context.Context) (dto.MeterWithoutService, error)
}

type meter struct {
	repo repo.Repository
}

func NewMeterApp(meterRepo repo.Repository) Meter {
	return &meter{meterRepo}
}

func (app *meter) RegisterNewMeter(ctx context.Context, request dto.Meter) error {
	reqModel := model.Meter{}
	reqModel.BuildModel(request)

	meterDB, err := app.GetMeterByBrandAndSerial(ctx, request.Brand, request.Serial)
	if err != nil {
		return err
	}

	if meterDB.ID.ID() != 0 {
		return echo.NewHTTPError(http.StatusConflict, "this meter brand and serial already exist")
	}

	if err := app.repo.RegisterNewMeter(ctx, reqModel); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (app *meter) GetMeterByBrandAndSerial(ctx context.Context, brand, serial string) (model.Meter, error) {
	existingMeter, err := app.repo.GetMeterByBrandAndSerial(ctx, brand, serial)
	if err != nil {
		return model.Meter{}, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return existingMeter, nil
}

func (app *meter) GetMeterByID(ctx context.Context, ID uuid.UUID) (model.Meter, error) {
	meterDB, err := app.repo.GetMeterByID(ctx, ID)
	if err != nil {
		return model.Meter{}, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if meterDB.ID.ID() == 0 {
		return model.Meter{}, echo.NewHTTPError(http.StatusBadRequest, "this meter does not exist")
	}

	return meterDB, nil
}

func (app *meter) DeleterMeter(ctx context.Context, ID uuid.UUID) error {
	meterDB, err := app.GetMeterByID(ctx, ID)
	if err != nil {
		return err
	}

	if meterDB.InUse {
		return echo.NewHTTPError(http.StatusConflict, "this meter is currently begin used")
	}

	if err = app.repo.DeleteMeterByID(ctx, ID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (app *meter) GetInactiveServiceMeters(ctx context.Context) (dto.MeterWithoutService, error) {
	clientMeter, err := app.repo.GetInactiveServiceMeters(ctx)
	if err != nil {
		return dto.MeterWithoutService{}, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return clientMeter.ToDomainDTOSlice(), nil
}
