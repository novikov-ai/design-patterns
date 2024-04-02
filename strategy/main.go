package main

func main() {
	rocket := rocket{}
	rocketDuck := New(rocket)

	rocketDuck.Fly()

	// change strategy from rocket to walker
	walker := walker{}
	rocketDuck.flyer = walker

	rocketDuck.Fly()
}
