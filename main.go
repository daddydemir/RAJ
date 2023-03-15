package main

import (
	"net/http"

	"github.com/daddydemir/RAJ/handlers"
)

func main() {
	println("server started at http://localhost:2023/api/v1/create/model")

	server := &http.Server{
		Addr:    ":2023",
		Handler: handlers.UrlDefine(),
	}
	server.ListenAndServe()
}
