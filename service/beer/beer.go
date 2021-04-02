package beer

import (
	"github.com/igiannoulas/go-sqlite-gpibdvhc/service/randomapi"
)

// RandomBeer d
type RandomBeer struct {
	resource string
	randomapi.RandomAPI
}

// GetRandomBeer d
func (beer RandomBeer) GetRandomBeer() string {
	return string(beer.GetItemFunc("beer/random_beer"))
}
