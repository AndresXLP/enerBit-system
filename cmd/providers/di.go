package providers

import (
	"enerBit-system/internal/app"
	pgRepo "enerBit-system/internal/infra/adapters/postgres/repo"
	redisRepo "enerBit-system/internal/infra/adapters/redis/repo"
	"enerBit-system/internal/infra/api/handler"
	"enerBit-system/internal/infra/api/router"
	"enerBit-system/internal/infra/api/router/groups"
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

	_ = Container.Provide(groups.NewMeterGroup)

	_ = Container.Provide(handler.NewMeterHandler)

	_ = Container.Provide(app.NewMeterApp)

	_ = Container.Provide(pgRepo.NewMeterRepository)

	_ = Container.Provide(redisRepo.NewRedisRepository)

	return Container
}
