package server

import (
	"github.com/codegangsta/negroni"
	"github.com/tahmazidik/Go-microservice/api/router"
	"github.com/tahmazidik/Go-microservice/internal/config"
	"net/http"
)

// Server - структура сервера
type Server struct {
	AppConfig *config.Config
	Router    *router.Router
}

// NewServer - Создание нового сервера
func NewServer(appConfig *config.Config, router *router.Router) *Server {
	return &Server{
		AppConfig: appConfig,
		Router:    router,
	}
}

// Run - метод для запуска сервера
func (server *Server) Run() {
	ngRouter := server.Router.InitRoutes()
	ngClassic := negroni.Classic()
	ngClassic.UseHandler(ngRouter)
	err := http.ListenAndServe(":5000", ngClassic)
	if err != nil {
		return
	}

}
