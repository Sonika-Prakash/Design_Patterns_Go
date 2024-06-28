package main

// VictorianChair is the concrete chair
type VictorianChair struct {
	Variant string
	Price   int
}

func (a *VictorianChair) getVariant() string {
	return a.Variant
}

func (a *VictorianChair) getPrice() int {
	return a.Price
}
