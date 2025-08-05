package main

type IChair interface {
	setLegs(legs int)
	getLegs() int
	getvariant() string
}

type Chair struct {
	variant string
	legs    int
}

func (c *Chair) setLegs(legs int) {
	c.legs = legs
}

func (c *Chair) getLegs() int {
	return c.legs
}

func (c *Chair) getvariant() string {
	return c.variant
}
