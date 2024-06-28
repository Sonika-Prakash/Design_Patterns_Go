package main

const (
	artDecoSofaPrice   = 1000
	victorianSofaPrice = 1500
	modernSofaPrice    = 1200
)

// ISofa is the abstract chair interface
type ISofa interface {
	getVariant() string
	getPrice() int
}
