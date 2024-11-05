package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/razidev/movie-festival/src/configs"
	"github.com/razidev/movie-festival/src/routers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	configs.Connect()

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: routers.InitRoutes(),
	}

	server.ListenAndServe()
}
