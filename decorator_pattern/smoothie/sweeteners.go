package main

type Sweeteners struct {
	smoothie ISmoothie
}

func (s *Sweeteners) addIngedients(ingredients []string) {
	s.smoothie.addIngedients(ingredients)
}

func (s *Sweeteners) getIngredients() []string {
	return s.smoothie.getIngredients()
}
