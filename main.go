package main

import (
	"coffeemachine/blender"
	"coffeemachine/drinks"
	"coffeemachine/ingredientsstore"
	"coffeemachine/jsonParser"
	"coffeemachine/machine"
	"coffeemachine/outlet"
	"fmt"
)

const (
	filePath = "resources/coffeemachine.json"
)

func main() {
	dep := jsonParser.BuildDependencies(filePath)
	ingredsStore := ingredientsstore.New(dep.IngredientsVolumeCapacity)
	coffeemachine := machine.NewMachine(createOutlets(dep, ingredsStore), ingredsStore)

	for i := 0; i < 200; i++ {
		d, err := coffeemachine.Prepare(1, drinks.BLACK_TEA)
		if err != nil {
			fmt.Println("error occurred while preparing GreenTea:", err)
			return
		}

		fmt.Printf("Here is your Drink: %+v\n", d)
	}
}

func createOutlets(dep jsonParser.Dependency, ingredsStore drinks.Store) map[int]machine.Percolator {
	machineOutlets := make(map[int]machine.Percolator, dep.Outlets)

	for i := 0; i < dep.Outlets; i++ {
		machineOutlets[i+1] = buildOutlet(i+1, ingredsStore, dep.Drinks)
	}

	return machineOutlets
}

func buildOutlet(id int, ingredientStore drinks.Store, supportedDrinks map[drinks.DrinkType]drinks.Recipe) *outlet.Outlet {
	var blackTeaMaker blender.BlackTeaMaker
	var hotTeaMaker blender.HotTeaMaker
	var hotCoffeeMaker blender.HotCoffeeMaker
	var greenTeaMaker blender.GreenTeaMaker

	for drinkType, recipe := range supportedDrinks {
		switch drinkType {
		case drinks.BLACK_TEA:
			blackTeaMaker = drinks.NewBlackTeaMaker(recipe, ingredientStore)

		case drinks.GREEN_TEA:
			greenTeaMaker = drinks.NewGreenTeaMaker(recipe, ingredientStore)

		case drinks.HOT_TEA:
			hotTeaMaker = drinks.NewHotTeaMaker(recipe, ingredientStore)

		case drinks.HOT_COFFEE:
			hotCoffeeMaker = drinks.NewHotCoffeeMaker(recipe, ingredientStore)
		}
	}

	supportedDrinkTypes := make(map[drinks.DrinkType]bool, len(supportedDrinks))

	for k := range supportedDrinks {
		supportedDrinkTypes[k] = true
	}

	b := blender.New(hotCoffeeMaker, greenTeaMaker, blackTeaMaker, hotTeaMaker, supportedDrinkTypes)
	return outlet.NewOutlet(id, b)
}
