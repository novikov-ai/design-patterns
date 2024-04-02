package main

import "fmt"

type SubscriberThird struct {
}

func (sf SubscriberThird) Update() {
	fmt.Println("SubscriberThird got update!")
}
