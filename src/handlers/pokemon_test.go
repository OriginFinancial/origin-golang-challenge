package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPokemonPagination(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		res := func2(1)
		assert.Equal(t, "https://pokeapi.co/api/v2/pokemon?offset=0&limit=20", res)
	})

	t.Run("test2", func(t *testing.T) {
		res := func2(2)
		assert.Equal(t, "https://pokeapi.co/api/v2/pokemon?offset=20&limit=20", res)
	})
}
