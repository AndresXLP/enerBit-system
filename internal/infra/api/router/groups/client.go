package groups

import (
	"enerBit-system/internal/infra/api/handler"
	"github.com/labstack/echo/v4"
)

type Client interface {
	Resource(c *echo.Group)
}

type client struct {
	meterHandler handler.Client
}

func NewClientGroup(meterHand handler.Client) Client {
	return &client{meterHand}
}

func (group client) Resource(c *echo.Group) {
	groupPath := c.Group("/client")
	groupPath.POST("", group.meterHandler.NewInstallation)
	groupPath.DELETE("", group.meterHandler.UninstallMeter)
}
