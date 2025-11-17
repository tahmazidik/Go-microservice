package book

import (
	"fmt"
	"github.com/tahmazidik/Go-microservice/internal/entities/book"
)

// Repository - структура репозитория
type Repository struct {
	database []book.Entity
}

// NewRepository - Метод для регистрации в DI контейнере
func NewRepository() *Repository {
	return &Repository{
		database: make([]book.Entity, 0),
	}
}

// Create - метод для создания книги в репозитории
func (repository *Repository) Create(entity book.Entity) {
	repository.database = append(repository.database, entity)
	fmt.Println(repository.database)
}
