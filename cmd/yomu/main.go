package main

import (
	"flag"

	"github.com/umizu/yomu/internal/api"
)

func main() {
	listenAddr := flag.String("listenAddr", ":8000", "listen address of the API server")
	flag.Parse()
	
	server, err := api.NewAPIServer(*listenAddr)
	if err != nil {
		panic(err)
	}

	server.Run()
}
