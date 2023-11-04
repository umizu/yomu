package contracts

import (
	"errors"
)

var (
	ErrInvalidLibraryItemBookId = errors.New("invalid library item book id")
	ErrInvalidStatus            = errors.New("invalid status")
)

type CreateLibraryItemRequest struct {
	BookId string `json:"bookId"`
	Status string `json:"status"`
}

func (r CreateLibraryItemRequest) Validate() error {
	if r.BookId == "" {
		return ErrInvalidLibraryItemBookId
	}
	if r.Status != "Planning" && r.Status != "Reading" && r.Status != "Completed" && r.Status != "Dropped" {
		return ErrInvalidStatus
	}
	return nil
}