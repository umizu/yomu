package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/umizu/yomu/internal/handlers"
)

func (s *APIServer) RegisterBookRoutes() {
	s.router.GET("/books", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		handlers.BooksGETHandler(w, r, p, s.db)
	})
	
	s.router.POST("/books", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		handlers.BooksPOSTHandler(w, r, p, s.db)
	})
}
