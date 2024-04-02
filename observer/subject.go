package main

type Subject struct {
	observers []Observer
}

func New() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
	}
}

func (s *Subject) Subscribe(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) Unsubscribe(observer Observer) {
	for i, o := range s.observers {
		if observer == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *Subject) Notify() {
	for _, o := range s.observers {
		o.Update()
	}
}
