package app

import (
	"github.com/tahmazidik/Go-microservice/api/controllers/book"
	"github.com/tahmazidik/Go-microservice/api/router"
	"github.com/tahmazidik/Go-microservice/internal/config"
	bookRepositories "github.com/tahmazidik/Go-microservice/internal/repositories/book"
	bookService "github.com/tahmazidik/Go-microservice/internal/services/book"
	"github.com/tahmazidik/Go-microservice/server"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	_ = container.Provide(config.NewConfig)
	_ = container.Provide(server.NewServer)
	_ = container.Provide(router.NewRouter)

	buildBook(container)

	return container
}

func buildBook(container *dig.Container) {
	_ = container.Provide(book.NewController)
	_ = container.Provide(book.NewControllerRoute)
	_ = container.Provide(bookService.NewService)
	_ = container.Provide(bookRepositories.NewRepository)
}
