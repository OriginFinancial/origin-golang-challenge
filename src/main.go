package main

import (
	"log"
	"net/http"

	"github.com/originfinancial/origin-golang-challenge/src/handlers"
	"github.com/originfinancial/origin-golang-challenge/src/main/httputils"
)

func main() {
	env := "dev"
	port := "3000"

	log.Printf("🚀 Server starting on port %s", port)

	serverAddr := httputils.GetAddr(env, port)
	if err := http.ListenAndServe(serverAddr, httputils.RemoveTrailingSlashes(handlers.New())); err != nil {
		log.Fatalf("Exited with error %s", err.Error())
	}
}
