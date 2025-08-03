package main

// CheeseTopping is the concrete decorator for adding cheese to a pizza.
type CheeseTopping struct {
	pizza IPizza // field for referencing the wrapped object
}

func (c *CheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice() // get the price of the wrapped pizza
	return pizzaPrice + 30
}
