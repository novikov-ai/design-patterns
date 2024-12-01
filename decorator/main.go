package main

import "fmt"

func main() {
	beverage := Beverage{
		190, "Coffee", true, true, true,
	}

	// Without Decorator
	Checkout(beverage)

	// With Decorator
	beverageWithToppings := BeverageWithToppings{
		190, "Coffee", []Beverager{
			BeverageWithToppings{
				cost: 15, description: "milk",
			},
			BeverageWithToppings{
				cost: 30, description: "soya",
			},
			BeverageWithToppings{
				cost: 20, description: "chocolate",
			},
		},
	}

	Checkout(beverageWithToppings)
}
func Checkout(in Beverager) {
	fmt.Printf("%s will cost you: %d\n\n", in.Info(), in.Cost())
}
