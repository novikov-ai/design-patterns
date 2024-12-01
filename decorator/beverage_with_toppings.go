package main

import (
	"fmt"
	"strings"
)

// BeverageWithToppings - Decorator
type BeverageWithToppings struct {
	cost        int
	description string
	toppings    []Beverager
}

func (bwt BeverageWithToppings) Cost() int {
	cost := bwt.cost

	for _, top := range bwt.toppings{
		cost += top.Cost()
	}

	return cost
}

func (bwt BeverageWithToppings) Info() string {
	builder := strings.Builder{}
	builder.WriteString(bwt.description)

	if len(bwt.toppings) > 0 {
		builder.WriteString(" with:\n")
	}

	for _, top := range bwt.toppings{
		builder.WriteString(fmt.Sprintf("- %s\n", top.Info()))
	}

	return builder.String()
}