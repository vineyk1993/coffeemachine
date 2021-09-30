package drinks

type HotTeaMaker struct {
	recipe Recipe
	store  Store
}

func NewHotTeaMaker(recipe Recipe, ingredientStore Store) *HotTeaMaker {
	return &HotTeaMaker{
		recipe: recipe,
		store: ingredientStore,
	}
}

func(htm *HotTeaMaker) MakeHotTea() (*Drink, error) {
	_, err := htm.store.GetIngredients(htm.recipe.Ingredients)
	if err != nil {
		return nil, err
	}
	return NewDrink(HOT_TEA), nil
}