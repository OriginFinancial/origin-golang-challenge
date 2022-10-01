package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func New() http.Handler {
	router := mux.NewRouter()

	registerHelloHandler(router)

	return router
}
