package service

import (
	"fmt"

	"connectrpc.com/connect"
	dexv1 "github.com/rendaman0215/poke_api/schema/gen/go/dex/v1"
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
func GetImage(req *connect.Request[dexv1.GetImageRequest]) (*connect.Response[dexv1.GetImageResponse], error) {
	res := connect.NewResponse(&dexv1.GetImageResponse{
		Url: fmt.Sprintf(
			"https://s3-ap-northeast-1.amazonaws.com/pokedb.tokyo/sv/assets/pokemon/thumbs_128/pokemon-%04d-%02d.png",
			req.Msg.Id, req.Msg.Form,
		),
	})
	return res, nil
}
