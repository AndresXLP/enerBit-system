package app

import (
	"context"

	"enerBit-system/internal/domain/dto"
	"enerBit-system/internal/domain/ports/postgres/repo"
	"enerBit-system/internal/infra/adapters/postgres/model"
	"github.com/andresxlp/gosuite/errs"
	"github.com/google/uuid"
)

type Meter interface {
	RegisterNewMeter(ctx context.Context, request dto.Meter) error
	GetMeterByBrandAndSerial(ctx context.Context, brand, serial string) (model.Meter, error)
	GetMeterByID(ctx context.Context, ID uuid.UUID) (model.Meter, error)
	DeleterMeter(ctx context.Context, ID uuid.UUID) error
	GetInactiveServiceMeters(ctx context.Context) (dto.MeterWithoutService, error)
	GetLastInstallation(ctx context.Context, request dto.LastInstallation) (dto.Client, error)
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
		return errs.NewAppError(errs.ResourceDuplicated, "this meter brand and serial already exist")
	}

	if err = app.repo.RegisterNewMeter(ctx, reqModel); err != nil {
		return err
	}

	return nil
}

func (app *meter) GetMeterByBrandAndSerial(ctx context.Context, brand, serial string) (model.Meter, error) {
	existingMeter, err := app.repo.GetMeterByBrandAndSerial(ctx, brand, serial)
	if err != nil {
		return model.Meter{}, err
	}

	return existingMeter, nil
}

func (app *meter) GetMeterByID(ctx context.Context, ID uuid.UUID) (model.Meter, error) {
	meterDB, err := app.repo.GetMeterByID(ctx, ID)
	if err != nil {
		return model.Meter{}, err
	}

	if meterDB.ID.ID() == 0 {
		return model.Meter{}, errs.NewAppError(errs.ResourceNotFound, "this meter does not exist")
	}

	return meterDB, nil
}

func (app *meter) DeleterMeter(ctx context.Context, ID uuid.UUID) error {
	meterDB, err := app.GetMeterByID(ctx, ID)
	if err != nil {
		return err
	}

	if meterDB.InUse {
		return errs.NewAppError(errs.ResourceInUse, "this meter is currently begin used")
	}

	if err = app.repo.DeleteMeterByID(ctx, ID); err != nil {
		return err
	}

	return nil
}

func (app *meter) GetInactiveServiceMeters(ctx context.Context) (dto.MeterWithoutService, error) {
	clientMeter, err := app.repo.GetInactiveServiceMeters(ctx)
	if err != nil {
		return dto.MeterWithoutService{}, err
	}

	return clientMeter.ToDomainDTOSlice(), nil
}

func (app *meter) GetLastInstallation(ctx context.Context, request dto.LastInstallation) (dto.Client, error) {
	meterDB, err := app.GetMeterByBrandAndSerial(ctx, request.Brand, request.Serial)
	if err != nil {
		return dto.Client{}, err
	}

	if meterDB.ID.ID() == 0 {
		return dto.Client{}, errs.NewAppError(errs.ResourceNotFound, "this meter does not exist")
	}

	if meterDB.LastInstallation.IsZero() {
		return dto.Client{}, errs.NewAppError(errs.ResourceInvalid, "this meter has never been installed")
	}

	lastClient, err := app.repo.GetLastInstallation(ctx, meterDB)
	if err != nil {
		return dto.Client{}, err
	}

	return lastClient.ToDomainDTOSingle(), nil
}
