package main

// ModernFactory is the concrete factory
type ModernFactory struct{}

func (a *ModernFactory) createChair() IChair {
	return &ModernChair{
		Variant: modern,
		Price:   modernChairPrice,
	}
}

func (a *ModernFactory) createCoffeeTable() ICoffeeTable {
	return &ModernCoffeeTable{
		Variant: modern,
		Price:   modernCoffeeTablePrice,
	}
}

func (a *ModernFactory) createSofa() ISofa {
	return &ModernSofa{
		Variant: modern,
		Price:   modernSofaPrice,
	}
}
