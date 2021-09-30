package blender

import (
	"coffeemachine/drinks"
	"errors"
)

var (
	ErrDrinkNotSupported = errors.New("drink is not supported by the Blender")
)

// I could use Empty interface here i.e. Generic type supportedDrinks map[drinks.DrinkType]interface{}
// Which would mean a drink type has a corresponding maker; but that would require type casting here and there.
// So using concrete interfaces
type Blender struct {
	hcm HotCoffeeMaker
	gtm GreenTeaMaker
	btm BlackTeaMaker
	htm HotTeaMaker

	supportedDrinks map[drinks.DrinkType]bool
}

func New(hcm HotCoffeeMaker, gtm GreenTeaMaker, btm BlackTeaMaker, htm HotTeaMaker, supportedDrinks map[drinks.DrinkType]bool) *Blender {
	return &Blender{
		hcm: hcm,
		gtm: gtm,
		btm: btm,
		htm: htm,

		supportedDrinks: supportedDrinks,
	}
}

func (b *Blender) canBlend(drinkType drinks.DrinkType) bool {
	if _, ok := b.supportedDrinks[drinkType]; !ok {
		return false
	}
	return true
}

func (b *Blender) Blend(drinkType drinks.DrinkType) (*drinks.Drink, error) {
	if b.canBlend(drinkType) == false {
		return nil, ErrDrinkNotSupported
	}

	switch drinkType {
	case drinks.BLACK_TEA:
		if b.btm != nil {
			return b.btm.MakeBlackTea()
		}

	case drinks.GREEN_TEA:
		if b.gtm != nil {
			return b.gtm.MakeGreenTea()
		}

	case drinks.HOT_TEA:
		if b.htm != nil {
			return b.htm.MakeHotTea()
		}

	case drinks.HOT_COFFEE:
		if b.hcm != nil {
			return b.hcm.MakeHotCoffee()
		}
	}

	return nil, ErrDrinkNotSupported
}
