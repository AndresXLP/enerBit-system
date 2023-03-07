package router

import (
	"enerBit-system/internal/infra/api/handler"

	"github.com/labstack/echo/v4"
)

type Router struct {
	server *echo.Echo
}

func New(server *echo.Echo) *Router {
	return &Router{
		server,
	}
}

func (r *Router) Init() {
	basePath := r.server.Group("/api") //customize your basePath
	basePath.GET("/health", handler.HealthCheck)
}
