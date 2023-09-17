package api

import "github.com/umizu/yomu/internal/handlers"

func (s *APIServer) RegisterBookRoutes() {
	s.router.GET("/books", handlers.BooksGETHandler)
	s.router.POST("/books", handlers.BooksPOSTHandler)
}
