package handlers

type PokemonDTO struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonListResponse struct {
	Results []PokemonDTO `json:"results"`
}
