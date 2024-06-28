package main

// ArtDecoChair is the concrete chair
type ArtDecoChair struct {
	Variant string
	Price   int
}

func (a *ArtDecoChair) getVariant() string {
	return a.Variant
}

func (a *ArtDecoChair) getPrice() int {
	return a.Price
}
