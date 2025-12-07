package main

import "fmt"

type human struct {
	name string
	age  int
}

func newHuman(name string, age int) *human {
	return &human{
		name: name,
		age:  age,
	}
}

type Action struct {
	human
}

func (h human) printToString() {
	fmt.Printf("My name is %s. I have %d age", h.name, h.age)
}

func main() {
	m := newHuman("Marat", 22)
	a := Action{*m}
	a.printToString()
}
