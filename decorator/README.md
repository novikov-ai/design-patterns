# Decorator pattern

[BeverageWithToppings](beverage_with_toppings.go) is a decorator under [Beverage](beverage.go).

With the approach like that we can manage our code in runtime. 

## Just ask yourself: how many line of code I should change if I need to add an extra topping?

### Without decorator

- beverage.go -> Cost():
~~~go
func (b Beverage) Cost() int {
	// ...
    if b.NewTopping {
		cost += 100
	}
    // ...
}
~~~

- beverage.go -> Info():
~~~go
func (b Beverage) Info() string {
	// ...

	if b.hasMilk || b.hasSoya || b.hasChocolate || b.HasNewTopping{
		description += " with:\n"
	}

    // ...

	if b.NewTopping {
		description += fmt.Sprintf("- %s\n", "new topping")
	}

	// ...
}
~~~

- main.go -> main():
~~~go
func main() {
	beverage := Beverage{
		190, "Coffee", true, true, true, true
	}

    // ...
}
~~~

### With decorator

- main.go -> main():
~~~go
func main() {
    // ...
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
            BeverageWithToppings{
				cost: 100, description: "new topping",
			},
		},
	}
}
~~~