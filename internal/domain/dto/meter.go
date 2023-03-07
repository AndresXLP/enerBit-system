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
