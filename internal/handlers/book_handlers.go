package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/umizu/yomu/internal/data"
	"github.com/umizu/yomu/internal/models"
	"github.com/umizu/yomu/internal/util"
	"github.com/umizu/yomu/pkg/contracts"
)

func BooksGETHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params, db data.Store) {
	books, err := db.GetAllBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.WriteJSON(w, http.StatusOK, books)
}

func BooksPOSTHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params, db data.Store) {
	var request contracts.BookRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	book := models.NewBook(request.Title, request.MediaType, request.Length)

	err = db.CreateBook(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.WriteJSON(w, http.StatusCreated, book)
}
