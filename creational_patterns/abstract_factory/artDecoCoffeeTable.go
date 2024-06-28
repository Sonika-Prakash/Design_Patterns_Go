package main

// ArtDecoCoffeeTable  is the concrete chair
type ArtDecoCoffeeTable struct {
	Variant string
	Price   int
}

func (a *ArtDecoCoffeeTable) getVariant() string {
	return a.Variant
}

func (a *ArtDecoCoffeeTable) getPrice() int {
	return a.Price
}
