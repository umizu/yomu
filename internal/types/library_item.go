package types

import (
	"github.com/google/uuid"
	"github.com/umizu/yomu/pkg/contracts"
)

type LibraryItem struct {
	Id     string `json:"id"`
	BookId string `json:"bookId"`
	Status Status `json:"status"`
}

type Status int

const (
	Planning Status = iota
	Reading
	Completed
	Dropped
)

func NewLibraryItemFromRequest(req contracts.CreateLibraryItemRequest) *LibraryItem {
	return &LibraryItem{
		Id:   uuid.NewString(),
		BookId: req.BookId,
		Status: ParseStatus(req.Status),
	}
}

func ParseStatus(status string) Status {
	switch status {
	case "Planning":
		return Planning
	case "Reading":
		return Reading
	case "Completed":
		return Completed
	case "Dropped":
		return Dropped
	default:
		return Planning
	}
}

func (s Status) String() string {
	switch s {
	case Planning:
		return "Planning"
	case Reading:
		return "Reading"
	case Completed:
		return "Completed"
	case Dropped:
		return "Dropped"
	default:
		return ""
	}
}
