package main

// ModernCoffeeTable  is the concrete chair
type ModernCoffeeTable struct {
	Variant string
	Price   int
}

func (a *ModernCoffeeTable) getVariant() string {
	return a.Variant
}

func (a *ModernCoffeeTable) getPrice() int {
	return a.Price
}
