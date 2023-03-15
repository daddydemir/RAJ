package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func UrlDefine() http.Handler {
	r := mux.NewRouter().StrictSlash(true)
	baseUrl := "/api/v1/"

	r.HandleFunc(baseUrl+"create/model", test).Methods(POST)
	r.HandleFunc("/test", test).Methods(GET)

	handler := cors.AllowAll().Handler(r)
	return handler

}
