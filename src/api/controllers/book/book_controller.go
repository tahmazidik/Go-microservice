package book

import (
	"encoding/json"
	"net/http"

	bookModels "github.com/tahmazidik/Go-microservice/internal/models/book"
	"github.com/tahmazidik/Go-microservice/internal/services/book"
)

// Controller - для работы с книгами
type Controller struct {
	service *book.Service
}

// NewController - метод для регистрации в DI контейнере
func NewController(service *book.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (controller *Controller) CreateBook(w http.ResponseWriter, r *http.Request) {
	request := new(bookModels.CreateModel)
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(request)

	controller.service.Create(request)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
