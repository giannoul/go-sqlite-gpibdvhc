package main

import (
	"fmt"

	"github.com/igiannoulas/go-sqlite-gpibdvhc/service/beer"
)

func main() {
	mybeer := beer.RandomBeer{}
	fmt.Println(mybeer.GetRandomBeer())
}
