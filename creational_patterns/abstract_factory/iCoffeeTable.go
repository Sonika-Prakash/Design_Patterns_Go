package main

const (
	artDecoCoffeeTablePrice   = 400
	victorianCoffeeTablePrice = 600
	modernCoffeeTablePrice    = 650
)

// ICoffeeTable is the abstract chair interface
type ICoffeeTable interface {
	getVariant() string
	getPrice() int
}
