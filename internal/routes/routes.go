package routes

import (
	"net/http"
	"github.com/umizu/yomu/internal/handlers"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/books", handlers.BookHandler)
	return mux
}
