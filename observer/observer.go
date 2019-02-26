package main

import "fmt"

type Subject interface {
	AddObserver(observer Observer)
	RemoveObserver(observer Observer)
	Notify(str string)
}

type Observer interface {
	Update(str string)
}

type SubjectObj struct {
	state     string
	observers []Observer
}

func (this *SubjectObj) setState(state string) {
	this.state = state
	this.Notify(state)
}

func (this *SubjectObj) AddObserver(observer Observer) {
	this.observers = append(this.observers, observer)
}

func (this *SubjectObj) RemoveObserver(observer Observer) {
	for i, v := range this.observers {
		if v == observer {
			this.observers = append(this.observers[:i], this.observers[i+1:]...)
		}
	}
}

func (this *SubjectObj) Notify(str string) {
	for _, v := range this.observers {
		v.Update(str)
	}
}

type ObserverA struct{}

func (this ObserverA) Update(str string) {
	fmt.Println("I'm observer A, subject has changed to " + str)
}

type ObserverB struct{}

func (this ObserverB) Update(str string) {
	fmt.Println("I'm observer B,subject has changed to " + str)
}

type ObserverC struct{}

func (this ObserverC) Update(str string) {
	fmt.Println("I'm observer C,subject has changed to " + str)
}

func main() {
	sub := SubjectObj{}
	ob1, ob2, ob3 := ObserverA{}, ObserverB{}, ObserverC{}
	sub.AddObserver(ob1)
	sub.AddObserver(ob2)
	sub.AddObserver(ob3)
	sub.RemoveObserver(ob2)
	sub.setState("sleep")
}
