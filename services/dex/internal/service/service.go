package service

import (
	"fmt"

	"github.com/rendaman0215/poke_api/services/dex/internal/domain/models"
)

// PokemonThumbURL returns the S3 thumbnail URL for a Pokémon.
//
// id   : Pokédex ID (e.g. 1 for Bulbasaur)
// form : Form number (e.g. 0 for default form)
//
// Example:
//
//	PokemonThumbURL(25, 3)
//	=> "https://s3-ap-northeast-1.amazonaws.com/pokedb.tokyo/sv/assets/pokemon/thumbs_128/pokemon-0025-03.png"
func GetImage(pokeId models.PokemonId) (string, error) {
	res := fmt.Sprintf(
		"https://s3-ap-northeast-1.amazonaws.com/pokedb.tokyo/sv/assets/pokemon/thumbs_128/pokemon-%04d-%02d.png",
		pokeId.ID, pokeId.Form,
	)
	return res, nil
}
