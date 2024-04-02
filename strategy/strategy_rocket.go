package main

import (
	"fmt"
)

type rocket struct {
}

func (r rocket) Fly() {
	fmt.Println("fly like a rocket!")
}
