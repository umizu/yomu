package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/umizu/yomu/internal/models"
	"github.com/umizu/yomu/internal/util"
	"github.com/umizu/yomu/pkg/contracts"
)

func BooksGETHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "GET /books")
}

func BooksPOSTHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var request contracts.BookRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("title: %s\nmediaType: %s\nlength: %d",
		request.Title,
		request.MediaType,
		request.Length)

	book := models.NewBook(request.Title, request.MediaType, request.Length)

	util.WriteJSON(w, http.StatusCreated, book)
}
