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

func (s *APIServer) RegisterLibraryItemRoutes() {
	libraryItemHandler := handlers.NewLibraryItemHandler(data.NewPostgresLibraryItemStore(s.db))
	s.router.GET("/libraryItems", libraryItemHandler.LibraryItemGETHandler)
	s.router.POST("/libraryItems", libraryItemHandler.LibraryItemPOSTHandler)
}
