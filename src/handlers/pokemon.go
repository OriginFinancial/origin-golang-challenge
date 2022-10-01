package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type pokemonHandler struct{}

func registerPokemonHandler(r *mux.Router) {
	handler := pokemonHandler{}
	r = r.NewRoute().PathPrefix("/pokemon").Subrouter()

	r.HandleFunc("", handler.listPokemon).Methods("GET")
}

func (*pokemonHandler) listPokemon(w http.ResponseWriter, r *http.Request) {
	var err error
	var pageInt int

	page := r.URL.Query().Get("page")
	if page == "" {
		pageInt = 1
	} else {
		pageInt, err = strconv.Atoi(page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	pokemonsList, err := func1(pageInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pokemonsList)
}

func func1(page int) ([]PokemonDTO, error) {
	url := func2(page)

	client := http.Client{}
	r1, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	r1.Header = http.Header{
		"Authorization": {"eyJhbGciOiJIUzI1NiJ9.eyJSb2xlIjoiQWRtaW4iLCJJc3N1ZXIiOiJJc3N1ZXIiLCJVc2VybmFtZSI6IkphdmFJblVzZSIsImV4cCI6MTY2MjA0MjMzNCwiaWF0IjoxNjYyMDQyMzM0fQ.xi3uKpbHXXxE5iTOkDrkHJfpXQhGQGjLHXwC1SE-kFI"},
	}

	r2, err := client.Do(r1)
	if err != nil {
		return nil, err
	}
	defer r2.Body.Close()

	var r3 PokemonListResponse
	err = json.NewDecoder(r2.Body).Decode(&r3)
	if err != nil {
		return nil, err
	}

	return r3.Results, nil
}

func func2(page int) string {
	limit := 20
	offset := page - 1
	return fmt.Sprintf("https://pokeapi.co/api/v2/pokemon?offset=%d&limit=%d", offset, limit)
}
