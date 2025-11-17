package book

import "github.com/google/uuid"

// Entity - можель в БД для книги
type Entity struct {
	Uuid uuid.UUID
	Name string
}
