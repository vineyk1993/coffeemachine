package drinks

import "coffeemachine/ingredientsstore"

type Recipe struct {
	Instructions []string // will be executed in sequence
	Ingredients  map[ingredientsstore.Ingredient]int
}
