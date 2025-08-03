package main

type CornTopping struct {
	pizza IPizza
}

func (c *CornTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 15
}
