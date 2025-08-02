package main

// TomatoTopping is the concrete decorator for adding tomato to a pizza.
type TomatoTopping struct {
	pizza IPizza // field for referencing the wrapped object
}

func (t *TomatoTopping) getPrice() int {
	pizzaPrice := t.pizza.getPrice() // get the price of the wrapped pizza
	return pizzaPrice + 20
}
