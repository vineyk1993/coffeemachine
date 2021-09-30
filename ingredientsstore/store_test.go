package ingredientsstore

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStore_GetIngredients(t *testing.T) {
	storeIngreds := make(map[Ingredient]int)
	storeIngreds[Ingredient{Name: "hot_water"}] = 100
	storeIngreds[Ingredient{Name: "green_tea_powder"}] = 150

	store := New(storeIngreds)

	t.Run("should return IngredientsNotAvailable error", func (t *testing.T) {

		testIngreds := make(map[Ingredient]int)
		testIngreds[Ingredient{Name: "water"}] = 10
		_, err := store.GetIngredients(testIngreds)
		assert.EqualError(t, err, IngredientsNotAvailable.Error())
	})

	t.Run("should return InvalidIngredientsVolume error", func (t *testing.T) {

		testIngreds := make(map[Ingredient]int)
		testIngreds[Ingredient{Name: "hot_water"}] = -10
		_, err := store.GetIngredients(testIngreds)
		assert.EqualError(t, err, InvalidIngredientsVolume.Error())
	})

	t.Run("should return InsufficientIngredients error", func (t *testing.T) {

		testIngreds := make(map[Ingredient]int)
		testIngreds[Ingredient{Name: "hot_water"}] = 1000
		_, err := store.GetIngredients(testIngreds)
		assert.EqualError(t, err, InsufficientIngredients.Error())
	})

	t.Run("should fulfil the requirement", func (t *testing.T) {

		testIngreds := make(map[Ingredient]int)
		testIngreds[Ingredient{Name: "hot_water"}] = 50
		_, err := store.GetIngredients(testIngreds)
		assert.NoError(t, err)
		assert.Equal(t, store.getVolume(Ingredient{Name: "hot_water"}), 50)
	})
}

func TestStore_Refill(t *testing.T) {
	storeIngreds := make(map[Ingredient]int)
	storeIngreds[Ingredient{Name: "hot_water"}] = 100
	storeIngreds[Ingredient{Name: "green_tea_powder"}] = 150

	store := New(storeIngreds)

	t.Run("should return IngredientsNotAvailable error", func (t *testing.T) {

		testIngreds := make(map[Ingredient]int)
		testIngreds[Ingredient{Name: "water"}] = 10
		err := store.Refill(testIngreds)
		assert.EqualError(t, err, IngredientsNotAvailable.Error())
	})

	t.Run("should return InvalidIngredientsVolume error", func (t *testing.T) {

		testIngreds := make(map[Ingredient]int)
		testIngreds[Ingredient{Name: "hot_water"}] = -10
		err := store.Refill(testIngreds)
		assert.EqualError(t, err, InvalidIngredientsVolume.Error())
	})

	t.Run("should refill the ingredients", func (t *testing.T) {
		testIngreds := make(map[Ingredient]int)
		testIngreds[Ingredient{Name: "hot_water"}] = 50

		_, err := store.GetIngredients(testIngreds)
		assert.NoError(t, err)

		err = store.Refill(testIngreds)
		assert.NoError(t, err)
		assert.Equal(t, store.getVolume(Ingredient{Name: "hot_water"}), 100)
	})

	t.Run("should refill the ingredients if volume is given more that capacity", func (t *testing.T) {
		testIngreds := make(map[Ingredient]int)
		testIngreds[Ingredient{Name: "hot_water"}] = 50

		_, err := store.GetIngredients(testIngreds)
		assert.NoError(t, err)

		testIngreds[Ingredient{Name: "hot_water"}] = 500
		err = store.Refill(testIngreds)
		assert.NoError(t, err)
		assert.Equal(t, store.getVolume(Ingredient{Name: "hot_water"}), 100)
	})
}