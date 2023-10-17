package api

import (
	"database/sql"

	"github.com/umizu/yomu/internal/data"
)

func (s *APIServer) RegisterBookRoutes(db *sql.DB) {

	bookHandler := NewBookHandler(data.NewPostgresBookStore(db))
	s.router.GET("/books", bookHandler.BooksGETHandler)
	s.router.POST("/books", bookHandler.BooksPOSTHandler)
}
