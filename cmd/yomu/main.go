package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/umizu/yomu/internal/routes"
)

func main() {
	router := routes.NewRouter()

	port := 8080
	addr := fmt.Sprintf(":%d", port)

	fmt.Printf("Server listening on http://localhost%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
