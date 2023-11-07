package models

type Status int

const (
	Planning Status = iota
	Reading
	Completed
	Dropped
)
