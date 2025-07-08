package usuario

import "pokemonGame/pokemon"

type Usuario struct {
	Name            string
	ActivePokemon   []pokemon.Pokemon
	CapturedPokemon []pokemon.Pokemon
}
