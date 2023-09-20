package models

import (
	"github.com/google/uuid"
)

type Book struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	MediaType string `json:"mediaType"`
	Length    int    `json:"length"`
}

func NewBook(title string, mediaType string, length int) *Book {
	return &Book{
		Id:        uuid.NewString(),
		Title:     title,
		MediaType: mediaType,
		Length:    length,
	}
}
