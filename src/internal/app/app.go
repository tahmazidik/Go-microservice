package app

import (
	"github.com/tahmazidik/Go-microservice/api/router"
	"github.com/tahmazidik/Go-microservice/internal/config"
	"github.com/tahmazidik/Go-microservice/server"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	_ = container.Provide(config.NewConfig)
	_ = container.Provide(router.NewRouter)
	_ = container.Provide(server.NewServer)

	return container
}
