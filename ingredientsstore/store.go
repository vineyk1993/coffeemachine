package ingredientsstore

import (
	"errors"
	"sync"
)

var (
	InsufficientIngredients  = errors.New("insufficient ingredients")
	InvalidIngredientsVolume = errors.New("invalid ingredients volume")
	IngredientsNotAvailable  = errors.New("ingredients not available")
)

type Ingredient struct {
	Name string
}

type Store struct {
	sync.Mutex
	capacity                   map[Ingredient]int
	availableIngredientsVolume map[Ingredient]int
}

func New(ingredients map[Ingredient]int) *Store {
	available := make(map[Ingredient]int)
	for k, v := range ingredients {
		available[k] = v
	}

	return &Store{
		capacity:                   ingredients,
		availableIngredientsVolume: available,
	}
}

func (s *Store) getVolume(ingred Ingredient) int {
	s.Lock()
	defer s.Unlock()

	return s.availableIngredientsVolume[ingred]
}

func (s *Store) validateIngredients(incomingIngredients map[Ingredient]int) (bool, error) {
	for iIngred, iVolume := range incomingIngredients {
		availableVolume, found := s.availableIngredientsVolume[iIngred]
		if found == false {
			return false, IngredientsNotAvailable
		}

		if iVolume < 0 {
			return false, InvalidIngredientsVolume
		}

		if iVolume > availableVolume {
			return false, InsufficientIngredients
		}
	}

	return true, nil
}

func (s *Store) fulfill(incomingIngredients map[Ingredient]int) {
	for iIngred, iVolume := range incomingIngredients {
		s.availableIngredientsVolume[iIngred] -= iVolume
	}
}

// Reserves the Volume when called Atomically
func (s *Store) GetIngredients(incomingIngredients map[Ingredient]int) (bool, error) {
	s.Lock()
	defer s.Unlock()

	_, err := s.validateIngredients(incomingIngredients)
	if err != nil {
		return false, err
	}

	s.fulfill(incomingIngredients)
	return true, nil
}

// Refill the Ingredients till capacity reaches
// throws error in case -ve Ingredient Volume is specified
func (s *Store) validateRefillingIngredients(incomingIngredients map[Ingredient]int) (bool, error) {
	for iIngred, iVolume := range incomingIngredients {
		_, found := s.capacity[iIngred]
		if found == false {
			return false, IngredientsNotAvailable
		}

		if iVolume < 0 {
			return false, InvalidIngredientsVolume
		}

		// Do not want to fail here; will just take as much as needed.
		//if iVolume > maxCapacity {
		//	return false, InsufficientIngredients
		//}
	}

	return true, nil
}

func (s *Store) refill(incomingIngredients map[Ingredient]int) {
	for iIngred, iVolume := range incomingIngredients {
		capacityLeft := s.capacity[iIngred] - s.availableIngredientsVolume[iIngred]

		if iVolume > capacityLeft { // Capacity full
			s.availableIngredientsVolume[iIngred] = s.capacity[iIngred]

		} else { // Taking whatever is coming; since space is more that whats coming
			s.availableIngredientsVolume[iIngred] += iVolume
		}
	}
}

func (s *Store) Refill(refillingVolume map[Ingredient]int) error {
	s.Lock()
	defer s.Unlock()

	_, err := s.validateRefillingIngredients(refillingVolume)
	if err != nil {
		return err
	}

	s.refill(refillingVolume)
	return nil
}
