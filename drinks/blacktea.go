package drinks

type BlackTeaMaker struct {
	recipe Recipe
	store  Store
}

func NewBlackTeaMaker(recipe Recipe, ingredientStore Store) *BlackTeaMaker {
	return &BlackTeaMaker{
		recipe: recipe,
		store: ingredientStore,
	}
}

func (btm *BlackTeaMaker) MakeBlackTea() (*Drink, error) {
	_, err := btm.store.GetIngredients(btm.recipe.Ingredients)
	if err != nil {
		return nil, err
	}
	return NewDrink(BLACK_TEA), nil
}
