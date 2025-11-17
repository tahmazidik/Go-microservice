package book

import (
	"github.com/google/uuid"
	bookEntities "github.com/tahmazidik/Go-microservice/internal/entities/book"
	bookService "github.com/tahmazidik/Go-microservice/internal/models/book"
	"github.com/tahmazidik/Go-microservice/internal/repositories/book"
)

// Service - для работы с книгами
type Service struct {
	repository *book.Repository
}

// NewService - метод для регистрации в DI контейнере
func NewService(repository *book.Repository) *Service {
	return &Service{repository: repository}
}

// Create - метод для создания книги
func (service *Service) Create(model *bookService.CreateModel) {
	// Создаем сущность
	bookEntity := bookEntities.Entity{
		Uuid: uuid.New(),
		Name: model.Name,
	}

	service.repository.Create(bookEntity)
}
