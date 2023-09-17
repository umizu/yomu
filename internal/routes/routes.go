package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/umizu/yomu/internal/handlers"
)

func NewRouter() http.Handler {
	router := httprouter.New()

	router.GET("/books", handlers.GetBooksHandler)
	return router
}
