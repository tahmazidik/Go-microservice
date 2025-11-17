package router

import (
	"github.com/gorilla/mux"
	"github.com/tahmazidik/Go-microservice/api/controllers/book"
)

// Router - структура описывающие маршрутизатор конроллеров
type Router struct {
	BookRoutes *book.ControllerRoute
}

// NewRouter - создание нового маршрутизатора
func NewRouter(bookRoutes *book.ControllerRoute) *Router {
	return &Router{
		BookRoutes: bookRoutes,
	}
}

// InitRouters - инициализация маршрутизации API
func (routes *Router) InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = routes.BookRoutes.Route(router)
	return router
}
