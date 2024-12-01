package main

import "fmt"

type Beverage struct {
	cost         int
	description  string
	hasMilk      bool
	hasSoya      bool
	hasChocolate bool
}

func (b Beverage) Cost() int {
	cost := b.cost

	if b.hasMilk {
		cost += 15
	}
	if b.hasSoya {
		cost += 30
	}
	if b.hasChocolate {
		cost += 20
	}

	return cost
}

func (b Beverage) Info() string {
	description := b.description

	if b.hasMilk || b.hasSoya || b.hasChocolate {
		description += " with:\n"
	}

	if b.hasMilk {
		description += fmt.Sprintf("- %s\n", "milk")
	}
	if b.hasSoya {
		description += fmt.Sprintf("- %s\n", "soya")
	}
	if b.hasChocolate {
		description += fmt.Sprintf("- %s\n", "chocolate")
	}

	return description
}
