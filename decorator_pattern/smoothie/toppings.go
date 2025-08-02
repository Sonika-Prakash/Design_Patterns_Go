package main

type Toppings struct {
	smoothie ISmoothie
}

func (t *Toppings) addIngedients(ingredients []string) {
	t.smoothie.addIngedients(ingredients)
}

func (t *Toppings) getIngredients() []string {
	return t.smoothie.getIngredients()
}
