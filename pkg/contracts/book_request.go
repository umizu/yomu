package contracts

import (
	"errors"
)

var (
	ErrInvalidBookTitle    = errors.New("invalid book title")
	ErrInvalidBookIsbn     = errors.New("invalid book isbn")
	ErrInvalidBookFormat   = errors.New("invalid book format")
	ErrInvalidBookLink     = errors.New("invalid book link")
	ErrInvalidBookLanguage = errors.New("invalid book language")
)

type CreateBookRequest struct {
	Title    string `json:"title"`
	Isbn     string `json:"isbn"`
	Format   string `json:"format"`
	Link     string `json:"link"`
	Language string `json:"language"`
}

func (r CreateBookRequest) Validate() error {
	if r.Title == "" {
		return ErrInvalidBookTitle
	}
	if r.Isbn == "" {
		return ErrInvalidBookIsbn
	}
	if r.Format == "" {
		return ErrInvalidBookFormat
	}
	if r.Link == "" {
		return ErrInvalidBookLink
	}
	if r.Language == "" {
		return ErrInvalidBookLanguage
	}
	return nil
}
