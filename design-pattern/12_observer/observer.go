package observer

import "fmt"

type ISubject interface {
	Register(IObserver)
	Remove(IObserver)
	Notify(string)
}

type IObserver interface {
	Update(string)
}

type Subject struct {
	observers []IObserver
}

func (sub *Subject) Register(observer IObserver) {
	sub.observers = append(sub.observers, observer)
}

func (sub *Subject) Remove(observer IObserver) {
	for i, ob := range sub.observers {
		if ob == observer {
			sub.observers = append(sub.observers[:i], sub.observers[i+1:]...)
		}
	}
}

func (sub *Subject) Notify(msg string) {
	for _, ob := range sub.observers {
		ob.Update(msg)
	}
}

type Observer1 struct{}

func (Observer1) Update(msg string) {
	fmt.Printf("Observer1 msg: %s\n", msg)
}

type Observer2 struct{}

func (Observer2) Update(msg string) {
	fmt.Printf("Observer2 msg: %s\n", msg)
}
