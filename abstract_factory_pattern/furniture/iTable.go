package main

type ITable interface {
	setLegs(legs int)
	getLegs() int
	getvariant() string
}

type Table struct {
	variant string
	legs    int
}

func (t *Table) setLegs(legs int) {
	t.legs = legs
}

func (t *Table) getLegs() int {
	return t.legs
}

func (t *Table) getvariant() string {
	return t.variant
}
