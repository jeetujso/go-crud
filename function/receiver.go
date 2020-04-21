package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type agent struct {
	person
	kill bool
}

func (a agent) speak() {
	fmt.Println("receiver function", a.first, a.last)
}

func main() {
	p1 := person{
		first: "Jitendra",
		last:  "Jha",
		age:   28,
	}

	p2 := agent{
		person: p1,
		kill:   true,
	}
	p2.speak()
}
