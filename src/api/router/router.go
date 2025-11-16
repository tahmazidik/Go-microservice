package router

import (
	"github.com/gorilla/mux"
)

// Router - структура описывающие маршрутизатор конроллеров
type Router struct {
}

// NewRouter - создание нового маршрутизатора
func NewRouter() *Router {
	return &Router{}
}

// InitRouters - инициализация маршрутизации API
func (routes *Router) InitRoutes() *mux.Router {
	router := mux.NewRouter()
	return router
}
