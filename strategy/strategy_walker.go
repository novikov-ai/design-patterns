package main

import (
	"fmt"
)

type walker struct {
}

func (w walker) Fly() {
	fmt.Println("can't fly ;(")
}
