package repo

import (
	"context"
	"net/http"

	"enerBit-system/internal/domain/ports/postgres/repo"
	redis "enerBit-system/internal/domain/ports/redis/repo"
	"enerBit-system/internal/infra/adapters/postgres/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const (
	tableClients = "clients"
	tableMeters  = "meters"
)

type meter struct {
	db        *gorm.DB
	redisLogs redis.Repository
}

func NewMeterRepository(db *gorm.DB, redisLogs redis.Repository) repo.Repository {
	return meter{
		db,
		redisLogs,
	}
}

func (repo meter) RegisterNewMeter(ctx context.Context, reqModel model.Meter) error {
	err := repo.db.WithContext(ctx).
		Table(tableMeters).
		Create(&reqModel).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	go repo.redisLogs.SendStreamLog(reqModel.GenerateMessageCreatedMeter())

	return nil
}

func (repo meter) GetMeterByBrandAndSerial(ctx context.Context, brand, serial string) (model.Meter, error) {
	meterDB := model.Meter{}
	err := repo.db.WithContext(ctx).
		Table(tableMeters).
		Where("brand = ? AND serial = ?", brand, serial).
		Scan(&meterDB).Error
	if err != nil {
		return model.Meter{}, err
	}

	return meterDB, nil
}

func (repo meter) GetInstallationByAddress(ctx context.Context, address string) (model.Client, error) {
	clientDB := model.Client{}
	err := repo.db.WithContext(ctx).
		Table(tableClients).
		Where("address = ?", address).
		Order("installation_date DESC").
		Scan(&clientDB).
		Error
	if err != nil {
		return model.Client{}, err
	}

	return clientDB, nil
}

func (repo meter) NewInstallation(ctx context.Context, installation model.NewInstallation) error {
	err := repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).
			Table(tableClients).
			Create(&installation.Client).
			Error; err != nil {
			tx.Rollback()
			return err
		}

		installation.Meter.InUse = true

		if err := tx.WithContext(ctx).
			Table(tableMeters).
			Updates(installation.Meter).
			Error; err != nil {
			tx.Rollback()
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	go repo.redisLogs.SendStreamLog(installation.GenerateMessageInstallation())

	return nil
}

func (repo meter) UninstallMeter(ctx context.Context, property model.Client) error {
	err := repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).
			Table(tableClients).
			Omit("installation_date", "meter_id", "address", "id").
			Updates(property).
			Error; err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.WithContext(ctx).
			Table(tableMeters).
			Where("id = ?", property.MeterID).
			Update("in_use", false).
			Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (repo meter) GetMeterByID(ctx context.Context, ID uuid.UUID) (model.Meter, error) {
	meterDB := model.Meter{}
	err := repo.db.WithContext(ctx).
		Table(tableMeters).
		Where("id = ?", ID).
		Scan(&meterDB).Error
	if err != nil {
		return model.Meter{}, err
	}

	return meterDB, nil
}

func (repo meter) DeleteMeterByID(ctx context.Context, ID uuid.UUID) error {
	err := repo.db.WithContext(ctx).
		Table(tableMeters).
		Delete(model.Meter{}, "id=?", ID).
		Error
	if err != nil {
		return err
	}

	return nil
}

func (repo meter) GetInactiveServiceMeters(ctx context.Context) (model.ClientWithoutService, error) {
	clientMeter := model.ClientWithoutService{}
	err := repo.db.WithContext(ctx).
		Table(tableClients + " as c").
		Select("c.address, c.installation_date, m.brand,m.serial").
		Where("c.is_active = FALSE").
		Joins("left join meters as m on c.meter_id = m.id").
		Scan(&clientMeter).Error
	if err != nil {
		return model.ClientWithoutService{}, err
	}

	return clientMeter, nil
}
