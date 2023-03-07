package repo

import (
	"context"

	"enerBit-system/internal/infra/adapters/postgres/model"
)

type Repository interface {
	RegisterNewMeter(ctx context.Context, reqModel model.Meter) error
	GetMeterByBrandAndSerial(ctx context.Context, brand, serial string) (model.Meter, error)
}
