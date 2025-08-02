package main

type Base struct {
	smoothie ISmoothie
}

func (b *Base) addIngedients(ingredients []string) {
	b.smoothie.addIngedients(ingredients)
}

func (b *Base) getIngredients() []string {
	return b.smoothie.getIngredients()
}
