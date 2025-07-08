package main

import (
	"fmt"
	"pokemonGame/pokemon"
	"pokemonGame/regiao"
	"pokemonGame/usuario"
)

func main() {
	//pikachu := p.Pokemon{
	//	Name: "Pikachu",
	//	Type: "Electric",
	//	ID:   1,

	// fmt.Println("\nPokemon name is: ", pikachu.Name,
	//	"\n It's type is: ", pikachu.Type,
	//	"\n And it's ID is: ", pikachu.ID)

	//squirtle := p.Pokemon{
	//	Name: "Squirtle",
	//	Type: "Water",
	//	ID:   2,

	//fmt.Println("\nPokemon name is: ", squirtle.Name,
	//	"\n It's type is: ", squirtle.Type,
	//	"\n And it's ID is: ", squirtle.ID)

	// pokemonInfo := pikachu.GetInfo()
	// fmt.Println("Pokemon's Info: ", pokemonInfo)
	// pokemonInfo2 := squirtle.GetInfo()
	// fmt.Println("Pokemon's Info: ", pokemonInfo2)

	Pokemons := regiao.Pokemons
	for _, Pokemon := range Pokemons {
		fmt.Println(Pokemon.GetInfo())
	}

	user := usuario.Usuario{
		Name:          "Ash",
		ActivePokemon: []pokemon.Pokemon{},
	}
	user.ActivePokemon = append(user.ActivePokemon, Pokemons[0])
	fmt.Println(user)
	user.ActivePokemon[0] = Pokemons[1]
	fmt.Println(user)

	atkResult := user.ActivePokemon[0].PokeAttack(&Pokemons[2])
	fmt.Println(atkResult)
	capResult := user.ActivePokemon[0].Capture(&Pokemons[2])
	fmt.Println(capResult)
}
