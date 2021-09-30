package outlet

import (
	"coffeemachine/drinks"
)

type Blender interface {
	Blend(drinkType drinks.DrinkType) (*drinks.Drink, error)
}

type Outlet struct {
	id      int
	blender Blender
}

func NewOutlet(id int, blender Blender) *Outlet {
	return &Outlet{
		id:      id,
		blender: blender,
	}
}

func (o *Outlet) PrepareDrink(drinkType drinks.DrinkType) (*drinks.Drink, error) {
	return o.blender.Blend(drinkType)
}
