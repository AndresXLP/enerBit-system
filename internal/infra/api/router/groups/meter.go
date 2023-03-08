package groups

import (
	"enerBit-system/internal/infra/api/handler"
	"github.com/labstack/echo/v4"
)

type Meter interface {
	Resource(c *echo.Group)
}

type meter struct {
	meterHandler handler.Meter
}

func NewMeterGroup(meterHand handler.Meter) Meter {
	return &meter{meterHand}
}

func (group meter) Resource(c *echo.Group) {
	groupPath := c.Group("/meter")
	groupPath.POST("", group.meterHandler.RegisterNewMeter)
	groupPath.DELETE("/:id", group.meterHandler.DeleteMeter)
	groupPath.GET("/inactive", group.meterHandler.GetInactiveServiceMeters)
	groupPath.GET("/last-installation", group.meterHandler.GetLastInstallation)
}
