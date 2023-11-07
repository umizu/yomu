package models

type Activity struct {
	Id     string `json:"id"`
	Date   string `json:"date"`
	BookId string `json:"bookId"`
	Status Status `json:"status"`
}
