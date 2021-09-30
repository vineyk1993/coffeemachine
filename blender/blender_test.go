package blender

import (
	"coffeemachine/drinks"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockDrinkMaker struct {}

var makeHotTea func() (*drinks.Drink, error)
var makeGreenTea func() (*drinks.Drink, error)
var makeBlackTea func() (*drinks.Drink, error)
var makeHotCoffee func() (*drinks.Drink, error)

func (mdm *MockDrinkMaker) MakeHotTea() (*drinks.Drink, error) {
	return makeHotTea()
}

func (mdm *MockDrinkMaker) MakeGreenTea() (*drinks.Drink, error) {
	return makeGreenTea()
}

func (mdm *MockDrinkMaker) MakeBlackTea() (*drinks.Drink, error) {
	return makeBlackTea()
}

func (mdm *MockDrinkMaker) MakeHotCoffee() (*drinks.Drink, error) {
	return makeHotCoffee()
}

func TestBlend(t *testing.T) {
	t.Run("return ErrDrinkNotSupported if drink type is not supported", func(t *testing.T) {
		mockDrinkMaker := &MockDrinkMaker{}
		supportedDrinks := make(map[drinks.DrinkType]bool)
		supportedDrinks[drinks.BLACK_TEA] = true

		blender := New(mockDrinkMaker, mockDrinkMaker, mockDrinkMaker, mockDrinkMaker, supportedDrinks)
		_, err := blender.Blend(drinks.HOT_COFFEE)
		assert.Equal(t, err, ErrDrinkNotSupported)
	})

	t.Run("return ErrDrinkNotSupported if drink is supported but drink maker is damaged", func(t *testing.T) {
		mockDrinkMaker := &MockDrinkMaker{}
		supportedDrinks := make(map[drinks.DrinkType]bool)
		supportedDrinks[drinks.BLACK_TEA] = true

		blender := New(mockDrinkMaker, mockDrinkMaker, nil, mockDrinkMaker, supportedDrinks)
		_, err := blender.Blend(drinks.BLACK_TEA)
		assert.Equal(t, err, ErrDrinkNotSupported)
	})

	t.Run("return error if blackTea drinkMaker returns error", func(t *testing.T) {
		mockDrinkMaker := &MockDrinkMaker{}
		supportedDrinks := make(map[drinks.DrinkType]bool)
		supportedDrinks[drinks.BLACK_TEA] = true

		makeBlackTea = func() (*drinks.Drink, error) {
			return nil, errors.New("some error occurred while making drink")
		}

		blender := New(nil, nil, mockDrinkMaker, nil, supportedDrinks)
		_, err := blender.Blend(drinks.BLACK_TEA)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "some error occurred while making drink")
	})

	t.Run("return error if grenTea drinkMaker returns error", func(t *testing.T) {
		mockDrinkMaker := &MockDrinkMaker{}
		supportedDrinks := make(map[drinks.DrinkType]bool)
		supportedDrinks[drinks.GREEN_TEA] = true

		makeGreenTea = func() (*drinks.Drink, error) {
			return nil, errors.New("some error occurred while making green drink")
		}

		blender := New(nil, mockDrinkMaker, nil, nil, supportedDrinks)
		_, err := blender.Blend(drinks.GREEN_TEA)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "some error occurred while making green drink")
	})

	t.Run("return error if hotTea drinkMaker returns error", func(t *testing.T) {
		mockDrinkMaker := &MockDrinkMaker{}
		supportedDrinks := make(map[drinks.DrinkType]bool)
		supportedDrinks[drinks.HOT_TEA] = true

		makeHotTea = func() (*drinks.Drink, error) {
			return nil, errors.New("some error occurred while making hottea drink")
		}

		b := New(nil, nil, nil, mockDrinkMaker, supportedDrinks)
		_, err := b.Blend(drinks.HOT_TEA)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "some error occurred while making hottea drink")
	})

	t.Run("return error if hotCoffee drinkMaker returns error", func(t *testing.T) {
		mockDrinkMaker := &MockDrinkMaker{}
		supportedDrinks := make(map[drinks.DrinkType]bool)
		supportedDrinks[drinks.HOT_COFFEE] = true

		makeHotCoffee = func() (*drinks.Drink, error) {
			return nil, errors.New("some error occurred while making hotCoffee drink")
		}

		blender := New(mockDrinkMaker, nil, nil, nil, supportedDrinks)
		_, err := blender.Blend(drinks.HOT_COFFEE)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "some error occurred while making hotCoffee drink")
	})
}