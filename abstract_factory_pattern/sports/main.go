package main

import "fmt"

func main() {
	brand := "adidas" // or "nike"
	factory := createFactory(brand)
	if factory == nil {
		fmt.Println("unknown brand")
		return
	}

	shoe := factory.createShoe()
	shirt := factory.createShirt()

	fmt.Printf("Shoe created with logo %s and size %d\n", shoe.getLogo(), shoe.getSize())
	fmt.Printf("Shirt created with logo %s and size %d\n", shirt.getLogo(), shirt.getSize())

	shoe.setSize(9)
	shirt.setSize(38)
	shirt.setLogo("ADIDAS")
	shoe.setLogo("ADIDAS")
	fmt.Printf("Updated Shoe size: %d\n", shoe.getSize())
	fmt.Printf("Updated Shirt size: %d\n", shirt.getSize())
	fmt.Printf("Updated Shoe logo: %s\n", shoe.getLogo())
	fmt.Printf("Updated Shirt logo: %s\n", shirt.getLogo())
}
