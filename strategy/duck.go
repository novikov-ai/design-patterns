package main

type Duck struct {
	flyer Flyer
}

func New(f Flyer) Duck {
	return Duck{
		flyer: f,
	}
}

func (d Duck) Fly() {
	d.flyer.Fly()
}
