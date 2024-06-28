package main

// VictorianCoffeeTable  is the concrete chair
type VictorianCoffeeTable struct {
	Variant string
	Price   int
}

func (a *VictorianCoffeeTable) getVariant() string {
	return a.Variant
}

func (a *VictorianCoffeeTable) getPrice() int {
	return a.Price
}
