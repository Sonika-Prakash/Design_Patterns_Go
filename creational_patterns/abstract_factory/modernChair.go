package main

// ModernChair is the concrete chair
type ModernChair struct {
	Variant string
	Price   int
}

func (a *ModernChair) getVariant() string {
	return a.Variant
}

func (a *ModernChair) getPrice() int {
	return a.Price
}
