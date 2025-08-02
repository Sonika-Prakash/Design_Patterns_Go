package main

import "fmt"

func main() {
	base := []string{"water", "ice", "milk"}
	fruits := []string{"banana", "pineapple", "strawberry"}
	sweeteners := []string{"honey"}
	toppings := []string{"roasted cashews", "chocolate syrup"}

	smoothie := &Smoothie{
		ingredients: []string{},
	}

	fmt.Println("Starting with the base for the juice", base)
	addBase := &Base{
		smoothie: smoothie,
	}
	addBase.addIngedients(base)

	fmt.Println("Adding fruits to the smoothie", fruits)
	addFruits := &Fruits{
		smoothie: addBase,
	}
	addFruits.addIngedients(fruits)

	fmt.Println("Adding sweeteners to the smoothie", sweeteners)
	addSweeteners := &Sweeteners{
		smoothie: addFruits,
	}
	addSweeteners.addIngedients(sweeteners)

	fmt.Println("Finishing with toppings for the smoothie", toppings)
	addToppings := &Toppings{
		smoothie: addSweeteners,
	}
	addToppings.addIngedients(toppings)

	finalIngredients := addToppings.getIngredients()
	fmt.Println("Final ingredients in the smoothie:", finalIngredients)
	fmt.Println("Enjoy your smoothie!")
}
