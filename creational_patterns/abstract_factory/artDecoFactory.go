package main

// ArtDecoFactory is the concrete factory
type ArtDecoFactory struct{}

func (a *ArtDecoFactory) createChair() IChair {
	return &ArtDecoChair{
		Variant: artDeco,
		Price:   artDecoChairPrice,
	}
}

func (a *ArtDecoFactory) createCoffeeTable() ICoffeeTable {
	return &ArtDecoCoffeeTable{
		Variant: artDeco,
		Price:   artDecoCoffeeTablePrice,
	}
}

func (a *ArtDecoFactory) createSofa() ISofa {
	return &ArtDecoSofa{
		Variant: artDeco,
		Price:   artDecoSofaPrice,
	}
}
