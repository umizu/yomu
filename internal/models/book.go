package models

import (
	"github.com/google/uuid"
)

type Book struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Pages int    `json:"pages"`
}

func NewBook(title string, pages int) *Book {
	return &Book{
		Id:    uuid.NewString(),
		Title: title,
		Pages: pages,
	}
}
