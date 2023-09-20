package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/umizu/yomu/internal/data"
)

type APIServer struct {
	listenAddr string
	router     *httprouter.Router
	db         data.Store
}

func NewAPIServer(listenAddr string) (*APIServer, error) {
	db, err := data.NewPostgresStore()
	if err != nil {
		return nil, err
	}

	err = db.Init()
	if err != nil {
		return nil, err
	}

	return &APIServer{
		listenAddr: listenAddr,
		router:     httprouter.New(),
		db:         db,
	}, nil
}

func (s *APIServer) Run() {
	s.RegisterBookRoutes()
	fmt.Printf("Server listening on http://localhost%s\n", s.listenAddr)
	log.Fatal(http.ListenAndServe(s.listenAddr, s.router))
}
