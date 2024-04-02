package main

import "fmt"

type SubscriberSecond struct {
}

func (sf SubscriberSecond) Update() {
	fmt.Println("SubscriberSecond got update!")
}
