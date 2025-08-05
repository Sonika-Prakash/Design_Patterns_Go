package main

type IFurnitureFactory interface {
	createChair() IChair
	createTable() ITable
}

func createFactory(variant string) IFurnitureFactory {
	switch variant {
	case "modern":
		return &ModernFurnitureFactory{}
	case "classic":
		return &ClassicFurnitureFactory{}
	default:
		return nil
	}
}
