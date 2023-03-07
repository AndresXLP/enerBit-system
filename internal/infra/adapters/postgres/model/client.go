package model

import (
	"time"

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
