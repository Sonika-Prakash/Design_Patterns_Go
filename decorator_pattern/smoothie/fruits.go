package main

type Fruits struct {
	smoothie ISmoothie
}

func (f *Fruits) addIngedients(ingredients []string) {
	f.smoothie.addIngedients(ingredients)
}

func (f *Fruits) getIngredients() []string {
	return f.smoothie.getIngredients()
}
