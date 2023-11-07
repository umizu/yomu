package models

import (
	"github.com/google/uuid"
	"github.com/umizu/yomu/pkg/contracts"
)

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Isbn     string `json:"isbn"`
	Format   string `json:"format"`
	Link     string `json:"link"`
	Language string `json:"language"`
}

func NewBookFromRequest(req contracts.CreateBookRequest) *Book {
	return &Book{
		ID:       uuid.NewString(),
		Title:    req.Title,
		Isbn:     req.Isbn,
		Format:   req.Format,
		Link:     req.Link,
		Language: req.Language,
	}
}
