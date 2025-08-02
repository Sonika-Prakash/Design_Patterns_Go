package main

type NonvegPizza struct{}

func (n *NonvegPizza) getPrice() int {
	return 120
}
