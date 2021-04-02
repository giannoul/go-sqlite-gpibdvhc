package main

import (
	"encoding/json"
	"fmt"

	beermodel "github.com/igiannoulas/go-sqlite-gpibdvhc/repository/model/beer"
	beersvc "github.com/igiannoulas/go-sqlite-gpibdvhc/service/beer"
)

func main() {
	mybeersvc := beersvc.RandomBeer{}
	mybeermodel := beermodel.Beer{}

	if err := json.Unmarshal(mybeersvc.GetRandomBeer(), &mybeermodel); err != nil {
		panic(err)
	}

	fmt.Println(mybeermodel)
	fmt.Printf("%#v", mybeermodel)
}
