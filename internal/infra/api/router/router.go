package router

import (
	"enerBit-system/internal/infra/api/handler"
	"enerBit-system/internal/infra/api/router/groups"

	"github.com/labstack/echo/v4"
)

type Router struct {
	server      *echo.Echo
	meterGroup  groups.Meter
	clientGroup groups.Client
}

func New(server *echo.Echo, meterGroup groups.Meter, clientGroup groups.Client) *Router {
	return &Router{
		server,
		meterGroup,
		clientGroup,
	}
}

func (r *Router) Init() {
	basePath := r.server.Group("/api") //customize your basePath
	basePath.GET("/health", handler.HealthCheck)

	r.meterGroup.Resource(basePath)
	r.clientGroup.Resource(basePath)
}
