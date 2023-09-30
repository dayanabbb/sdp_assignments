package main

import (
	"fmt"
	"sync"
)

type Observer interface {
	Update(newDeadline string)
}

type Subject interface {
	AddObserver(observer Observer)
	NotifyObservers()
	SetNewDeadline(newDeadline string)
}

type Assignment struct {
	observers     []Observer
	deadline      string
	observersLock sync.Mutex
}

func NewAssignment(oldDeadline string) *Assignment {
	return &Assignment{
		deadline: oldDeadline,
	}
}

func (a *Assignment) AddObserver(observer Observer) {
	a.observersLock.Lock()
	defer a.observersLock.Unlock()
	a.observers = append(a.observers, observer)
}

func (a *Assignment) NotifyObservers() {
	a.observersLock.Lock()
	defer a.observersLock.Unlock()
	for _, observer := range a.observers {
		observer.Update(a.deadline)
	}
}

func (a *Assignment) SetNewDeadline(newDeadline string) {
	a.deadline = newDeadline
	a.NotifyObservers()
}

type Student struct {
	name string
}

func NewStudent(name string) *Student {
	return &Student{name}
}

func (s *Student) Update(newDeadline string) {
	fmt.Println("student %s received new deadline: &s\n", s.name, newDeadline)
}

func main() {

	//new assignment and students
	assignment := NewAssignment("2023-09-30")
	student1 := NewStudent("Alissa")
	student2 := NewStudent("Dayana")

	//added students
	assignment.AddObserver(student1)
	assignment.AddObserver(student2)

	//changed the deadline
	assignment.SetNewDeadline("2023-10-01")

}
