package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type helloHandler struct{}

func registerHelloHandler(r *mux.Router) {
	handler := helloHandler{}
	r = r.NewRoute().PathPrefix("/hello").Subrouter()

	r.HandleFunc("/world", handler.world).Methods("GET")
}

func (*helloHandler) world(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(
		HelloWorldResponse{Message: "Hello World"},
	)
}
