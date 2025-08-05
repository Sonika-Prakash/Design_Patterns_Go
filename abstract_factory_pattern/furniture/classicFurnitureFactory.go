package main

type ClassicFurnitureFactory struct{}

func (c *ClassicFurnitureFactory) createChair() IChair {
	return &ClassicChair{
		Chair: Chair{
			variant: "classic",
			legs:    4, // Default number of legs for a classic chair
		},
	}
}

func (c *ClassicFurnitureFactory) createTable() ITable {
	return &ClassicTable{
		Table: Table{
			variant: "classic",
			legs:    4, // Default number of legs for a classic table
		},
	}
}
