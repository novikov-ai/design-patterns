package main

import "fmt"

func main() {
	subject := New()

	sf := SubscriberFirst{}
	subject.Subscribe(sf)

	ss := SubscriberSecond{}
	subject.Subscribe(ss)

	st := SubscriberThird{}
	subject.Subscribe(st)

	subject.Notify()

	fmt.Println("... SubscriberSecond unsubscribed... ")
	subject.Unsubscribe(ss)

	subject.Notify()
}
