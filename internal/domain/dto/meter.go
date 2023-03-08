package dto

import (
	"context"
	"time"

	"github.com/go-playground/mold/v4/modifiers"
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
	conform  = modifiers.New()
)

type Meter struct {
	Brand  string `json:"brand" mod:"ucase" validate:"required"`
	Serial string `json:"serial" mod:"ucase" validate:"required"`
	Lines  int    `json:"lines" validate:"min=1,max=10"`
}

func (m *Meter) Validate() error {
	_ = conform.Struct(context.Background(), m)
	return validate.Struct(m)
}

type NewInstallation struct {
	Meter
	Client
}

func (n *NewInstallation) Validate() error {
	_ = conform.Struct(context.Background(), n)

	return validate.Struct(n)
}

type UninstallMeter struct {
	Address        string    `json:"address" mod:"ucase" validate:"required"`
	RetirementDate time.Time `json:"retirement_date" validate:"required" example:"2023-03-10T00:00:00Z"`
}

func (u *UninstallMeter) Validate() error {
	_ = conform.Struct(context.Background(), u)

	return validate.Struct(u)
}

type ClientMeter struct {
	Address          string    `json:"address"`
	InstallationDate time.Time `json:"installation_date" example:"2023-03-10T00:00:00Z"`
	Brand            string    `json:"brand"`
	Serial           string    `json:"serial"`
	IsActive         bool      `json:"is_active"`
}

type MeterWithoutService []ClientMeter

func (m *MeterWithoutService) Add(clientMeter ClientMeter) {
	*m = append(*m, clientMeter)
}

type LastInstallation struct {
	Brand  string `query:"brand" mod:"ucase" validate:"required"`
	Serial string `query:"serial" mod:"ucase" validate:"required"`
}

func (l *LastInstallation) Validate() error {
	_ = conform.Struct(context.Background(), l)

	return validate.Struct(l)
}
