package drinks

import (
	"coffeemachine/ingredientsstore"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeGreeTea(t *testing.T) {
	t.Run("return error if store returns error", func(t *testing.T) {
		getIngredients = func(ingredients map[ingredientsstore.Ingredient]int) (bool, error) {
			return false, errors.New("some error occurred")
		}

		btm := NewHotCoffeeMaker(Recipe{}, &MockIStore{})
		_, err := btm.MakeHotCoffee()
		assert.EqualError(t, err, "some error occurred")
	})

	t.Run("return drink in case no error from store", func(t *testing.T) {
		getIngredients = func(ingredients map[ingredientsstore.Ingredient]int) (bool, error) {
			return true, nil
		}

		btm := NewHotCoffeeMaker(Recipe{}, &MockIStore{})
		drink, err := btm.MakeHotCoffee()
		assert.NoError(t, err)
		assert.Equal(t, drink, &Drink{Type: HOT_COFFEE})
	})
}