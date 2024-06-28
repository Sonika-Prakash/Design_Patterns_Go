package main

import "fmt"

const (
	artDeco   = "Art Deco"
	victorian = "Victorian"
	modern    = "Modern"
)

// IFurnitureFactory is the abstract factory interface
type IFurnitureFactory interface {
	createChair() IChair
	createCoffeeTable() ICoffeeTable
	createSofa() ISofa
}

func getFactoryVariant(variant string) (IFurnitureFactory, error) {
	if variant == artDeco {
		return &ArtDecoFactory{}, nil
	}
	if variant == victorian {
		return &VictorianFactory{}, nil
	}
	if variant == modern {
		return &ModernFactory{}, nil
	}

	return nil, fmt.Errorf("No such factory type")
}
