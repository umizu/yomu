package models

import (
	"github.com/google/uuid"
	"github.com/umizu/yomu/pkg/contracts"
)

type LibraryItem struct {
	Id     string `json:"id"`
	BookId string `json:"bookId"`
	Status Status `json:"status"`
}

func NewLibraryItemFromRequest(req contracts.UpsertLibraryItemRequest) *LibraryItem {
	return &LibraryItem{
		Id:     uuid.NewString(),
		BookId: req.BookId,
		Status: ParseStatus(req.Status),
	}
}

func (li *LibraryItem) ToResponse() contracts.LibraryItemResponse {
	return contracts.LibraryItemResponse{
		BookId: li.BookId,
		Status: li.Status.String(),
	}
}

func ToResponse(li []*LibraryItem) []contracts.LibraryItemResponse {
	resp := []contracts.LibraryItemResponse{}
	for _, item := range li {
		resp = append(resp, item.ToResponse())
	}
	return resp
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
