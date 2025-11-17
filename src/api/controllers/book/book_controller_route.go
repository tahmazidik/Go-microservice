package book

import "github.com/gorilla/mux"

// ControllerRoute настройки маршрутизации для нашего контроллера
type ControllerRoute struct {
	Controller *Controller
}

// NewControllerRoute метод для регистрации в DI контейнере
func NewControllerRoute(controller *Controller) *ControllerRoute {
	return &ControllerRoute{
		Controller: controller,
	}
}

// Route добавить в роутер маршрут
func (route *ControllerRoute) Route(router *mux.Router) *mux.Router {
	router.HandleFunc("/books", route.Controller.CreateBook).Methods("POST")
	return router
}
