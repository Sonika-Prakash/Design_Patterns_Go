package main

// VictorianSofa  is the concrete chair
type VictorianSofa struct {
	Variant string
	Price   int
}

func (a *VictorianSofa) getVariant() string {
	return a.Variant
}

func (a *VictorianSofa) getPrice() int {
	return a.Price
}
