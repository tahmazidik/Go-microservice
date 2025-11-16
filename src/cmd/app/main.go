package main

import (
	"github.com/tahmazidik/Go-microservice/internal/app"
	"github.com/tahmazidik/Go-microservice/internal/config"
	"github.com/tahmazidik/Go-microservice/server"
)

func main() {
	// Загрузка окружения
	config.LoadEnviroment()

	// Создаем DI контейнер
	container := app.BuildContainer()

	// Инвоким сервер
	err := container.Invoke(func(server *server.Server) {
		server.Run()
	})
	if err != nil {
		panic(err)
	}
}
