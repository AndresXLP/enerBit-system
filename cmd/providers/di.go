package providers

import (
	"enerBit-system/internal/infra/api/router"
	"enerBit-system/internal/infra/resource/postgres"
	"enerBit-system/internal/infra/resource/redis"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

var Container *dig.Container

func BuildContainer() *dig.Container {
	Container = dig.New()

	_ = Container.Provide(func() *echo.Echo {
		return echo.New()
	})

	_ = Container.Provide(redis.NewRedisConnection)

	_ = Container.Provide(postgres.NewConnection)

	_ = Container.Provide(router.New)

	return Container
}
