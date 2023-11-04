package api

import (
	"github.com/umizu/yomu/internal/data"
	"github.com/umizu/yomu/internal/handlers"
)

func (s *APIServer) RegisterBookRoutes() {
	bookHandler := handlers.NewBookHandler(data.NewPostgresBookStore(s.db))
	s.router.GET("/books", bookHandler.BooksGETHandler)
	s.router.POST("/books", bookHandler.BooksPOSTHandler)
}

func (s *APIServer) RegisterBookStatusRoutes() {
	bookStatusHandler := handlers.NewBookStatusHandler(data.NewPostgresBookStatusStore(s.db))
	s.router.GET("/bookStatuses", bookStatusHandler.BookStatusGETHandler)
	s.router.POST("/bookStatuses", bookStatusHandler.BookStatusPOSTHandler)
}
