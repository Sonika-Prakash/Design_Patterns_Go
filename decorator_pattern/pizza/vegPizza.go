package main

// VegPizza is the concrete component for the pizza decorator pattern.
type VegPizza struct{}

func (v *VegPizza) getPrice() int {
	return 80
}
