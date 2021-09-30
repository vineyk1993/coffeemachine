package jsonParser

import (
	"coffeemachine/drinks"
	"coffeemachine/ingredientsstore"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

func BuildDependencies(filePath string) Dependency {
	coffeeMachineJson, err := readJsonFile(filePath)
	if err != nil {
		panic("error while reading json file: " + err.Error())
	}

	outletsCount, err := getOutletsCount(coffeeMachineJson)
	if err != nil {
		panic(err)
	}

	totalIngredients, err := getTotalIngredientsVolume(coffeeMachineJson)
	if err != nil {
		panic(err)
	}

	supportedDrinks, err := getSupportedDrinksAndRecipies(coffeeMachineJson)
	if err != nil {
		panic(err)
	}

	deps := Dependency{
		Outlets:                   outletsCount,
		IngredientsVolumeCapacity: totalIngredients,
		Drinks:                    supportedDrinks,
	}

	return deps
}

func readJsonFile(filePath string) (map[string]interface{}, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var coffeeMachineBuilder map[string]interface{}
	err = json.Unmarshal(bytes, &coffeeMachineBuilder)
	if err != nil {
		return nil, err
	}

	coffeeMachineBuilder, ok := coffeeMachineBuilder["machine"].(map[string]interface{})
	if ok == false {
		return nil, errors.New("provided json is not valid")
	}

	return coffeeMachineBuilder, nil
}

func getOutletsCount(from map[string]interface{}) (int, error) {
	outlets, ok := from["outlets"].(map[string]interface{})
	if ok == false {
		return 0, errors.New("error while fetching outlets; provided json is not valid")
	}

	count, ok := outlets["count_n"].(float64)
	if ok == false {
		return 0, errors.New("error while fetching outlets; provided json is not valid")
	}

	return int(count), nil
}

func getTotalIngredientsVolume(from map[string]interface{}) (map[ingredientsstore.Ingredient]int, error) {
	totalIngredientsQuantity, ok := from["total_items_quantity"].(map[string]interface{})
	if ok == false {
		return nil, errors.New("total_items_quantity is not of expected type")
	}

	return getIngredients(totalIngredientsQuantity)
}

func getIngredients(from map[string]interface{}) (map[ingredientsstore.Ingredient]int, error) {
	ingredients := make(map[ingredientsstore.Ingredient]int)
	for k, v := range from {
		quantity, ok := v.(float64)
		if ok == false {
			return nil, errors.New("ingredient quantity should be integer")
		}

		ingredients[ingredientsstore.Ingredient{Name: k}] = int(quantity)
	}

	return ingredients, nil
}

func getSupportedDrinksAndRecipies(from map[string]interface{}) (map[drinks.DrinkType]drinks.Recipe, error) {
	beverages, ok := from["beverages"].(map[string]interface{})
	if ok == false {
		return nil, errors.New("error while fetching beverages; provided json is not valid")
	}

	supportedDrinks := make(map[drinks.DrinkType]drinks.Recipe)

	for k, v := range beverages {
		value, ok := v.(map[string]interface{})
		if ok == false {
			return nil, errors.New("error while fetching beverages; provided json is not valid")
		}

		switch k {
		case "hot_tea":
			drinkIngredients, err := getIngredients(value)
			if err != nil {
				return nil, err
			}
			supportedDrinks[drinks.HOT_TEA] = buildRecipe(drinkIngredients)

		case "hot_coffee":
			drinkIngredients, err := getIngredients(value)
			if err != nil {
				return nil, err
			}
			supportedDrinks[drinks.HOT_COFFEE] = buildRecipe(drinkIngredients)

		case "black_tea":
			drinkIngredients, err := getIngredients(value)
			if err != nil {
				return nil, err
			}
			supportedDrinks[drinks.BLACK_TEA] = buildRecipe(drinkIngredients)

		case "green_tea":
			drinkIngredients, err := getIngredients(value)
			if err != nil {
				return nil, err
			}
			supportedDrinks[drinks.GREEN_TEA] = buildRecipe(drinkIngredients)
		}
	}

	return supportedDrinks, nil
}

func buildRecipe(ingredients map[ingredientsstore.Ingredient]int) drinks.Recipe {
	instructions := make([]string, len(ingredients))

	i := 0
	for ing := range ingredients {
		instructions[i] = ing.Name
		i++
	}

	return drinks.Recipe{
		Instructions: instructions,
		Ingredients:  ingredients,
	}
}