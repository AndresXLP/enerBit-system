package model

import (
	"time"

	"enerBit-system/internal/domain/dto"
	"github.com/google/uuid"
)

type Client struct {
	ID               int
	Address          string
	IsActive         bool
	MeterID          uuid.UUID
	InstallationDate time.Time
	RetirementDate   *time.Time
}

func (c *Client) ToDomainDTOSingle() dto.Client {
	return dto.Client{
		Address:          c.Address,
		IsActive:         &c.IsActive,
		InstallationDate: c.InstallationDate,
	}
}
