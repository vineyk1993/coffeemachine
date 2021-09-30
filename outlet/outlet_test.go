package outlet

import (
	"coffeemachine/drinks"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockBlender struct {}
var blend func(drinkType drinks.DrinkType) (*drinks.Drink, error)

func (mb* MockBlender)Blend(drinkType drinks.DrinkType) (*drinks.Drink, error) {
	return blend(drinkType)
}

func TestPrepareDrink(t *testing.T) {
	t.Run("should return error in case blender not able to prepare drink", func(t *testing.T) {
		outlet := NewOutlet(1, &MockBlender{})

		blend = func(drinkType drinks.DrinkType) (*drinks.Drink, error) {
			return nil, errors.New("blender is not working")
		}

		_, err := outlet.PrepareDrink(drinks.HOT_COFFEE)
		assert.EqualError(t, err, "blender is not working")
	})

	t.Run("should return drink", func(t *testing.T) {
		outlet := NewOutlet(1, &MockBlender{})

		blend = func(drinkType drinks.DrinkType) (*drinks.Drink, error) {
			return &drinks.Drink{}, nil
		}

		d, err := outlet.PrepareDrink(drinks.HOT_COFFEE)
		assert.NoError(t, err)
		assert.Equal(t, d, &drinks.Drink{})
	})
}
