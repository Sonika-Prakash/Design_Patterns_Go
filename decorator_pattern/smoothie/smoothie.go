package main

type ISmoothie interface {
	addIngedients(ingredients []string)
	getIngredients() []string
}

type Smoothie struct {
	ingredients []string
}

func (s *Smoothie) addIngedients(ingredients []string) {
	s.ingredients = append(s.ingredients, ingredients...)
}

func (s *Smoothie) getIngredients() []string {
	return s.ingredients
}
