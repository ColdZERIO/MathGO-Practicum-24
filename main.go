package main

import (
	"fmt"
	"sort"
)

type Product struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
}

func main() {
	var Products = make(map[float64]Product)

	bread := Product{
		ID:       777,
		Name:     "Bread",
		Price:    50.5,
		Quantity: 3,
	}

	kolbasa := Product{
		ID:       222,
		Name:     "Kolbasa",
		Price:    2.2,
		Quantity: 2,
	}

	cheese := Product{
		ID:       111,
		Name:     "SyrSUKA",
		Price:    25.5,
		Quantity: 1,
	}

	prod := [3]Product{bread, kolbasa, cheese}

	prices := make([]float64, 0)

	for _, value := range prod {
		Products[value.Price] = value
		prices = append(prices, value.Price)
	}

	sort.Float64s(prices)

	for _, value := range prices {
		for _, v := range Products {
			if value == v.Price {
				fmt.Println(v.Name, v)
			}
		}
	}

}
