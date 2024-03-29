package main

import (
	"github.com/daddydemir/RAJ/config"
	"log"
	"net/http"

	"github.com/daddydemir/RAJ/internal/handlers"
)

func main() {
	log.Println("server started at http://localhost" + config.Get("PORT") + "/api/v1/create/model")

	server := &http.Server{
		Addr:    config.Get("PORT"),
		Handler: handlers.UrlDefine(),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
