package drinks

import (
	"coffeemachine/ingredientsstore"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockIStore struct{}

var getIngredients func(ingredients map[ingredientsstore.Ingredient]int) (bool, error)

func (mis *MockIStore) GetIngredients(ingredients map[ingredientsstore.Ingredient]int) (bool, error) {
	return getIngredients(ingredients)
}

func TestMakeBlackTea(t *testing.T) {
	t.Run("return error if store returns error", func(t *testing.T) {
		getIngredients = func(ingredients map[ingredientsstore.Ingredient]int) (bool, error) {
			return false, errors.New("some error occurred")
		}

		btm := NewBlackTeaMaker(Recipe{}, &MockIStore{})
		_, err := btm.MakeBlackTea()
		assert.EqualError(t, err, "some error occurred")
	})

	t.Run("return drink in case no error from store", func(t *testing.T) {
		getIngredients = func(ingredients map[ingredientsstore.Ingredient]int) (bool, error) {
			return true, nil
		}

		btm := NewBlackTeaMaker(Recipe{}, &MockIStore{})
		drink, err := btm.MakeBlackTea()
		assert.NoError(t, err)
		assert.Equal(t, drink, &Drink{Type: BLACK_TEA})
	})
}
