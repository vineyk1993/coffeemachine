package machine

import (
	"coffeemachine/drinks"
	"coffeemachine/ingredientsstore"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

// ======================================== Mock Outlet ============================================
type MockPercolator struct{}

var prepareDrink func(drinkType drinks.DrinkType) (*drinks.Drink, error)

func (mp *MockPercolator) PrepareDrink(drinkType drinks.DrinkType) (*drinks.Drink, error) {
	return prepareDrink(drinkType)
}

// ======================================== Mock Refiller Store ============================================
type MockIngredientsRefiller struct{}

var refillIngredients func(refillingVolume map[ingredientsstore.Ingredient]int) error

func (mir *MockIngredientsRefiller) Refill(refillingVolume map[ingredientsstore.Ingredient]int) error {
	return refillIngredients(refillingVolume)
}

func TestMachine_Prepare(t *testing.T) {
	t.Run("should return error in case invalid outlet provided", func(t *testing.T) {
		coffeeMachine := NewMachine(nil, nil)
		_, err := coffeeMachine.Prepare(0, drinks.HOT_COFFEE)
		assert.EqualError(t, err, InvalidPercolator.Error())
	})

	t.Run("should return requested drink", func(t *testing.T) {
		mp := &MockPercolator{}
		outlets := map[int]Percolator{1: mp}
		coffeeMachine := NewMachine(outlets, nil)

		prepareDrink = func(drinkType drinks.DrinkType) (*drinks.Drink, error) {
			return nil, errors.New("outlet not able to prepare drink")
		}

		_, err := coffeeMachine.Prepare(1, drinks.HOT_COFFEE)
		assert.EqualError(t, err, "outlet not able to prepare drink")
	})

	t.Run("should return requested drink if outlet could prepare requested drink", func(t *testing.T) {
		mp := &MockPercolator{}
		outlets := map[int]Percolator{1: mp}
		coffeeMachine := NewMachine(outlets, nil)

		prepareDrink = func(drinkType drinks.DrinkType) (*drinks.Drink, error) {
			return drinks.NewDrink(drinks.HOT_COFFEE), nil
		}

		d, err := coffeeMachine.Prepare(1, drinks.HOT_COFFEE)
		assert.NoError(t, err)
		assert.Equal(t, d, drinks.NewDrink(drinks.HOT_COFFEE))
	})
}

func TestMachine_RefillIngredients(t *testing.T) {
	t.Run("should return error in case refiller return error", func(t *testing.T) {
		coffeeMachine := NewMachine(nil, &MockIngredientsRefiller{})

		refillIngredients = func(refillingVolume map[ingredientsstore.Ingredient]int) error {
			return errors.New("could not refill the Ingredients")
		}

		err := coffeeMachine.RefillIngredients(nil)
		assert.EqualError(t, err, "could not refill the Ingredients")
	})

	t.Run("should return success if refiller succeeds", func(t *testing.T) {
		coffeeMachine := NewMachine(nil, &MockIngredientsRefiller{})

		refillIngredients = func(refillingVolume map[ingredientsstore.Ingredient]int) error {
			return nil
		}

		err := coffeeMachine.RefillIngredients(nil)
		assert.NoError(t, err)
	})
}
