package machine

import (
	"coffeemachine/drinks"
	"coffeemachine/ingredientsstore"
	"errors"
)

var (
	InvalidPercolator = errors.New("invalid percolator")
)

type CoffeeMachine interface {
	Prepare(outletID int, drinkType drinks.DrinkType) (*drinks.Drink, error)
	RefillIngredients(refillingVolume map[ingredientsstore.Ingredient]int) error
}

type Percolator interface {
	PrepareDrink(drinkType drinks.DrinkType) (*drinks.Drink, error)
}

type Refiller interface {
	Refill(refillingVolume map[ingredientsstore.Ingredient]int) error
}

type Machine struct {
	availableOutlets    map[int]Percolator
	ingredientsRefiller Refiller
}

func NewMachine(outlets map[int]Percolator, refiller Refiller) *Machine {
	return &Machine{
		availableOutlets:    outlets,
		ingredientsRefiller: refiller,
	}
}

func (m *Machine) Prepare(outletID int, drinkType drinks.DrinkType) (*drinks.Drink, error) {
	if _, ok := m.availableOutlets[outletID]; !ok {
		return nil, InvalidPercolator
	}

	return m.availableOutlets[outletID].PrepareDrink(drinkType)
}

func (m *Machine) RefillIngredients(refillingVolume map[ingredientsstore.Ingredient]int) error {
	return m.ingredientsRefiller.Refill(refillingVolume)
}
