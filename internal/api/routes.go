package api

import (
	"database/sql"

	"github.com/umizu/yomu/internal/data"
	"github.com/umizu/yomu/internal/handlers"
)

func (s *APIServer) RegisterBookRoutes(db *sql.DB) {

	bookHandler := handlers.NewBookHandler(data.NewPostgresBookStore(db))
	s.router.GET("/books", bookHandler.BooksGETHandler)
	s.router.POST("/books", bookHandler.BooksPOSTHandler)
}
