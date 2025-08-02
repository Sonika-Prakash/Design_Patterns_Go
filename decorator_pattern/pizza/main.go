package main

import "fmt"

func main() {
	pizza := &VegPizza{} // create a VegPizza instance
	// add cheese topping
	pizzaWithCheese := &CheeseTopping{
		pizza: pizza,
	}
	// add tomato topping
	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithCheese,
	}
	takeOrder(pizzaWithCheeseAndTomato)

	nvpizza := &NonvegPizza{} // create a NonvegPizza instance
	// add cheese topping
	pizzaWithCheese = &CheeseTopping{
		pizza: nvpizza,
	}
	pizzaWithCheeseAndCorn := &CornTopping{
		pizza: pizzaWithCheese,
	}
	takeOrder(pizzaWithCheeseAndCorn)
}

func takeOrder(pizza IPizza) {
	fmt.Println("Taking order for pizza with price:", pizza.getPrice())
}
