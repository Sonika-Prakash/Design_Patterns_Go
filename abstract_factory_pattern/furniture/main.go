package main

import "fmt"

func main() {
	factory := createFactory("modern") // or "classic"
	if factory == nil {
		fmt.Println("Unknown furniture factory variant")
		return
	}

	chair := factory.createChair()
	table := factory.createTable()
	fmt.Printf("Created a chair with %d legs\n", chair.getLegs())
	fmt.Printf("Created a table with %d legs\n", table.getLegs())

	chair.setLegs(3)
	table.setLegs(6)
	fmt.Printf("Updated chair to have %d legs\n", chair.getLegs())
	fmt.Printf("Updated table to have %d legs\n", table.getLegs())
	fmt.Printf("Chair variant: %s\n", chair.getvariant())
	fmt.Printf("Table variant: %s\n", table.getvariant())
}
