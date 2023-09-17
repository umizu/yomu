package main

import (
	"github.com/umizu/yomu/internal/api"
)

func main() {
	server := api.NewAPIServer(":8080")
	server.Run()
}
