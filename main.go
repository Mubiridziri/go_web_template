package main

import (
	"gin-template/models"
	"gin-template/routers"
	"net/http"
)

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: routers.InitRouter(),
	}

	models.ConnectDatabase()

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
