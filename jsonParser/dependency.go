package jsonParser

import (
	"coffeemachine/drinks"
	"coffeemachine/ingredientsstore"
)

type Dependency struct {
	Outlets                   int
	Drinks                    map[drinks.DrinkType]drinks.Recipe
	IngredientsVolumeCapacity map[ingredientsstore.Ingredient]int
}
