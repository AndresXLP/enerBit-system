package repo

import (
	"context"

	"enerBit-system/internal/infra/adapters/postgres/model"
	"github.com/google/uuid"
)

type Repository interface {
	RegisterNewMeter(ctx context.Context, reqModel model.Meter) error
	GetMeterByBrandAndSerial(ctx context.Context, brand, serial string) (model.Meter, error)
	GetInstallationByAddress(ctx context.Context, address string) (model.Client, error)
	NewInstallation(ctx context.Context, installation model.NewInstallation) error
	UninstallMeter(ctx context.Context, property model.Client) error
	GetMeterByID(ctx context.Context, ID uuid.UUID) (model.Meter, error)
	DeleteMeterByID(ctx context.Context, ID uuid.UUID) error
}
