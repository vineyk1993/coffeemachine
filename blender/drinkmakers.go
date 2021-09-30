package blender

import "coffeemachine/drinks"

type HotCoffeeMaker interface {
	MakeHotCoffee() (*drinks.Drink, error)
}

type GreenTeaMaker interface {
	MakeGreenTea() (*drinks.Drink, error)
}

type BlackTeaMaker interface {
	MakeBlackTea() (*drinks.Drink, error)
}

type HotTeaMaker interface {
	MakeHotTea() (*drinks.Drink, error)
}