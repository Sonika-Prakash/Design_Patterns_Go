package main

// ModernSofa  is the concrete chair
type ModernSofa struct {
	Variant string
	Price   int
}

func (a *ModernSofa) getVariant() string {
	return a.Variant
}

func (a *ModernSofa) getPrice() int {
	return a.Price
}
