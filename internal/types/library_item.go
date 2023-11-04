package types

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
