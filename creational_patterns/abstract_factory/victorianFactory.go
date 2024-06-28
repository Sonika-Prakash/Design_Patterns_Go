package main

// VictorianFactory is the concrete factory
type VictorianFactory struct{}

func (a *VictorianFactory) createChair() IChair {
	return &VictorianChair{
		Variant: victorian,
		Price:   victorianChairPrice,
	}
}

func (a *VictorianFactory) createCoffeeTable() ICoffeeTable {
	return &VictorianCoffeeTable{
		Variant: victorian,
		Price:   victorianCoffeeTablePrice,
	}
}

func (a *VictorianFactory) createSofa() ISofa {
	return &VictorianSofa{
		Variant: victorian,
		Price:   victorianSofaPrice,
	}
}
