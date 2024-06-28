package main

import "fmt"

func main() {
	modernFactory, _ := getFactoryVariant(modern)
	victorianFactory, _ := getFactoryVariant(victorian)
	artDecoFactory, _ := getFactoryVariant(artDeco)

	chair1 := modernFactory.createChair()
	fmt.Printf("Chair of type %s costs %d\n", chair1.getVariant(), chair1.getPrice())
	sofa1 := modernFactory.createSofa()
	fmt.Printf("Sofa of type %s costs %d\n", sofa1.getVariant(), sofa1.getPrice())
	coffeeTable1 := modernFactory.createCoffeeTable()
	fmt.Printf("Coffee Table of type %s costs %d\n\n", coffeeTable1.getVariant(), coffeeTable1.getPrice())

	chair2 := victorianFactory.createChair()
	fmt.Printf("Chair of type %s costs %d\n", chair2.getVariant(), chair2.getPrice())
	sofa2 := victorianFactory.createSofa()
	fmt.Printf("Sofa of type %s costs %d\n", sofa2.getVariant(), sofa2.getPrice())
	coffeeTable2 := victorianFactory.createCoffeeTable()
	fmt.Printf("Coffee Table of type %s costs %d\n\n", coffeeTable2.getVariant(), coffeeTable2.getPrice())

	chair3 := artDecoFactory.createChair()
	fmt.Printf("Chair of type %s costs %d\n", chair3.getVariant(), chair3.getPrice())
	sofa3 := artDecoFactory.createSofa()
	fmt.Printf("Sofa of type %s costs %d\n", sofa3.getVariant(), sofa3.getPrice())
	coffeeTable3 := artDecoFactory.createCoffeeTable()
	fmt.Printf("Coffee Table of type %s costs %d\n", coffeeTable3.getVariant(), coffeeTable3.getPrice())
}
