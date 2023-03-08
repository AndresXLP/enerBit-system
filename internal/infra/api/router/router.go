package router

import (
	_ "enerBit-system/docs"
	"enerBit-system/internal/infra/api/handler"
	"enerBit-system/internal/infra/api/router/groups"
	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
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
	r.server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status} latency=${latency_human}\n",
	}))
	r.server.Use(middleware.Recover())

	basePath := r.server.Group("/api") //customize your basePath
	basePath.GET("/health", handler.HealthCheck)
	basePath.GET("/swagger/*", echoSwagger.WrapHandler)

	r.meterGroup.Resource(basePath)
	r.clientGroup.Resource(basePath)
}
