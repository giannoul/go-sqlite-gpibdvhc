package beer

import (
	"github.com/igiannoulas/go-sqlite-gpibdvhc/service/randomapi"
)

// RandomBeer d
type RandomBeer struct {
	resource []byte
	randomapi.RandomAPI
}

// GetRandomBeer d
func (beer RandomBeer) GetRandomBeer() []byte {
	return beer.GetItemFunc("beer/random_beer")
}
