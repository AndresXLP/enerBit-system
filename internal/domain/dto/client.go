package dto

import "time"

type Client struct {
	Address          string    `json:"address" mod:"ucase" validate:"required"`
	IsActive         *bool     `json:"is_active" validate:"boolean"`
	InstallationDate time.Time `json:"installation_date" validate:"required" example:"2023-03-10T00:00:00Z"`
}
