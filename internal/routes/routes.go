package routes

import (
	"fmt"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/api/data", apiDataHandler)

	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home...")
}

func apiDataHandler(w http.ResponseWriter, r *http.Request) {
	data := "Some data..."
	fmt.Fprintln(w, data)
}
