package handlers

import (
	"fmt"
	"net/http"
)

func BookHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "GET /books")
	case http.MethodPost:
		fmt.Fprintf(w, "POST /books")
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
