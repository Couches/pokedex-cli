package main

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
  "fmt"

	"github.com/Couches/pokecache"
)

var pokemonCache pokecache.Cache = pokecache.NewCache(60 * time.Second)

func fetchPokemon(pokemonName string) (error, PokemonResponse) {
  pokemonResponse := PokemonResponse{}

  body, ok := pokemonCache.Get(pokemonName)

  if !ok {
    url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemonName)
    req, err := http.NewRequest("GET", url, nil)

    if err != nil {
      return err, pokemonResponse
    }

    res, err := http.DefaultClient.Do(req)

    if err != nil {
      return err, pokemonResponse
    }

    defer res.Body.Close()

    body, err = io.ReadAll(res.Body)

    if err != nil {
      return err, pokemonResponse
    }

    pokemonCache.Add(pokemonName, body)
  }

  err := json.Unmarshal(body, &pokemonResponse)

  if err != nil {
    return err, pokemonResponse
  }

  return nil, pokemonResponse
}
