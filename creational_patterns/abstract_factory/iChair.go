package main

const (
	artDecoChairPrice   = 100
	victorianChairPrice = 300
	modernChairPrice    = 250
)

// IChair is the abstract chair interface
type IChair interface {
	getVariant() string
	getPrice() int
}
