package main

import (
	"net/http"

	"github.com/razidev/movie-festival/src/configs"
	"github.com/razidev/movie-festival/src/routers"
)

func main() {
	configs.Connect()

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: routers.InitRoutes(),
	}

	server.ListenAndServe()
}
