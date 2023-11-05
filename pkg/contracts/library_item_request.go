package contracts

import (
	"errors"

	util "github.com/umizu/yomu/internal/utils"
)

var (
	ErrInvalidLibraryItemBookId = errors.New("invalid uuid format for bookId")
	ErrInvalidStatus            = errors.New("invalid status")
)

type CreateLibraryItemRequest struct {
	BookId string `json:"bookId"`
	Status string `json:"status"`
}

type LibraryItemResponse struct {
	Id     string `json:"id"`
	BookId string `json:"bookId"`
	Status string `json:"status"`
}

func (r CreateLibraryItemRequest) Validate() error {
	if !util.IsValidUUID(r.BookId) {
		return ErrInvalidLibraryItemBookId
	}
	if r.Status != "Planning" && r.Status != "Reading" && r.Status != "Completed" && r.Status != "Dropped" {
		return ErrInvalidStatus
	}
	return nil
}
