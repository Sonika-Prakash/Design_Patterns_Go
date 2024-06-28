package main

// ArtDecoSofa  is the concrete chair
type ArtDecoSofa struct {
	Variant string
	Price   int
}

func (a *ArtDecoSofa) getVariant() string {
	return a.Variant
}

func (a *ArtDecoSofa) getPrice() int {
	return a.Price
}
