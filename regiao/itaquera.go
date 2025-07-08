package regiao

import "pokemonGame/pokemon"

var Pokemons = []pokemon.Pokemon{
	{
		Name:     "Gengar",
		Type:     "Ghost",
		Attack:   50,
		Defense:  30,
		HP:       100,
		ID:       1,
		Captured: false,
	},

	{
		Name:     "Pikachu",
		Type:     "Electric",
		Attack:   60,
		Defense:  28,
		HP:       100,
		ID:       2,
		Captured: false,
	},

	{
		Name:     "Squirtle",
		Type:     "Water",
		Attack:   32,
		Defense:  40,
		HP:       100,
		ID:       3,
		Captured: false,
	},
}
