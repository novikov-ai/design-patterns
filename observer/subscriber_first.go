package main

import "fmt"

type SubscriberFirst struct {
}

func (sf SubscriberFirst) Update() {
	fmt.Println("SubscriberFirst got update!")
}
