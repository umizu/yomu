package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type APIServer struct {
	listenAddr string
	router     *httprouter.Router
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		router:     httprouter.New(),
	}
}

func (s *APIServer) Run() {
	s.RegisterBookRoutes()

	fmt.Printf("Server listening on http://localhost%s\n", s.listenAddr)
	log.Fatal(http.ListenAndServe(s.listenAddr, s.router))
}
