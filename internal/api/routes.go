package api

import (
	"github.com/umizu/yomu/internal/data"
	"github.com/umizu/yomu/internal/events"
	"github.com/umizu/yomu/internal/handlers"
)

func (s *APIServer) RegisterRoutes() {
	bookStore := data.NewPostgresBookStore(s.db)
	libraryItemStore := data.NewPostgresLibraryItemStore(s.db)

	messagech := make(chan interface{})
	bookHandler := handlers.NewBookHandler(bookStore)
	libraryItemHandler := handlers.NewLibraryItemHandler(libraryItemStore, bookStore, messagech)

	go events.Listen(messagech)

	s.router.GET("/books", bookHandler.BooksGETHandler)
	s.router.POST("/books", bookHandler.BooksPOSTHandler)
	s.router.GET("/libraryItems", libraryItemHandler.LibraryItemGETHandler)
	s.router.PUT("/libraryItems", libraryItemHandler.LibraryItemPUTHandler)
}
