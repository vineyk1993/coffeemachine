package drinks

type HotCoffeeMaker struct {
	recipe Recipe
	store  Store
}

func NewHotCoffeeMaker(recipe Recipe, ingredientStore Store) *HotCoffeeMaker {
	return &HotCoffeeMaker{
		recipe: recipe,
		store: ingredientStore,
	}
}

func(hcm *HotCoffeeMaker) MakeHotCoffee() (*Drink, error) {
	_, err := hcm.store.GetIngredients(hcm.recipe.Ingredients)
	if err != nil {
		return nil, err
	}
	return NewDrink(HOT_COFFEE), nil
}