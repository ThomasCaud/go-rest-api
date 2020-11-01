package model

import "github.com/google/uuid"

// Book struct is the model representation of a book.
type Book struct {
	ID    uuid.UUID `json:"uuid"`
	Title string    `json:"title"`
	Price int       `json:"price"`
}
