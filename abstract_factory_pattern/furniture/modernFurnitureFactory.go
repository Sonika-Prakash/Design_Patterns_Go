package main

type ModernFurnitureFactory struct{}

func (m *ModernFurnitureFactory) createChair() IChair {
	return &ModernChair{
		Chair: Chair{
			variant: "modern",
			legs:    4, // Default number of legs for a modern chair
		},
	}
}

func (m *ModernFurnitureFactory) createTable() ITable {
	return &ModernTable{
		Table: Table{
			variant: "modern",
			legs:    4, // Default number of legs for a modern table
		},
	}
}
