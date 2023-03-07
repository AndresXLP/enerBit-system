package model

import (
	"fmt"
	"time"

	"enerBit-system/internal/constants/enums"
	"enerBit-system/internal/domain/dto"
	"github.com/google/uuid"
)

type Meter struct {
	ID               uuid.UUID
	Brand            string
	Serial           string
	Lines            int
	InUse            bool
	LastInstallation time.Time
	CreatedAt        time.Time
}

func (m *Meter) BuildModel(meter dto.Meter) {
	m.ID = uuid.New()
	m.Brand = meter.Brand
	m.Serial = meter.Serial
	m.Lines = meter.Lines
}

func (m *Meter) GenerateMessageCreatedMeter() string {
	return fmt.Sprintf(enums.NewMeterRegister,
		m.Brand,
		m.Serial,
	)
}

type NewInstallation struct {
	Meter
	Client
}

func (m *NewInstallation) GenerateMessageInstallation() string {
	return fmt.Sprintf(enums.NewInstallation,
		m.MeterID,
		m.Address,
		m.InstallationDate.Day(),
		m.InstallationDate.Month().String(),
		m.InstallationDate.Year(),
		m.IsActive,
	)
}
