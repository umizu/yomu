package main

import (
	"github.com/umizu/yomu/internal/api"
)

func main() {
	server, err := api.NewAPIServer(":8080")
	if err != nil {
		panic(err)
	}

	server.Run()
}
