package drinks

import "coffeemachine/ingredientsstore"

type DrinkType string

const (
	BLACK_TEA  DrinkType = "BLACK_TEA"
	GREEN_TEA  DrinkType = "GREEN_TEA"
	HOT_TEA    DrinkType = "HOT_TEA"
	HOT_COFFEE DrinkType = "HOT_COFFEE"
)

type Drink struct {
	Type DrinkType
}

func NewDrink(drinkType DrinkType) *Drink {
	return &Drink{
		Type: drinkType,
	}
}

type Store interface {
	GetIngredients(ingredients map[ingredientsstore.Ingredient]int) (bool, error)
}