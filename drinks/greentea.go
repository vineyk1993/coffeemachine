package drinks

type GreenTeaMaker struct {
	recipe Recipe
	store  Store
}

func NewGreenTeaMaker(recipe Recipe, ingredientStore Store) *GreenTeaMaker {
	return &GreenTeaMaker{
		recipe: recipe,
		store: ingredientStore,
	}
}

func(gtm *GreenTeaMaker) MakeGreenTea() (*Drink, error) {
	_, err := gtm.store.GetIngredients(gtm.recipe.Ingredients)
	if err != nil {
		return nil, err
	}
	return NewDrink(GREEN_TEA), nil
}