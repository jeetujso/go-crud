package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type myBoss struct {
	person
	canFire bool
}

func main() {
	fmt.Println("Struct data type")
	//struct is a data structure which compose together the value of different types (aggregate together the value of dofferent type)
	p1 := person{
		first: "Jitendra",
		last:  "Jha",
		age:   28,
	}

	p2 := person{
		first: "Harpreet",
		last:  "Rupal",
		age:   32,
	}
	boss := myBoss{
		person: person{
			first: "Charan",
			last:  "Singh",
			age:   35,
		},
		canFire: true,
	}
	//var person []Person
	//fmt.Println(person) // []
	fmt.Println(p1) //{Jitendra Jha 28}
	fmt.Println(p2) // {Harpreet Rupal 32}

	// value of type person. it not an object. it is just a dot notation to get the value of type person
	fmt.Println(p1.first, p1.last, p1.age) // Jitendra
	fmt.Println(p2.first, p2.last, p2.age) //Rupal

	fmt.Println(boss)                                          //{{Charan Singh 35} true}
	fmt.Println(boss.first, boss.last, boss.age, boss.canFire) //Charan Singh 35 true. we can to get boss data like boss.persion.first

	// ======Anonymous struct======

	anonymousStruct := struct { //it is anonymous because it does has name
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "bond",
		age:   45,
	}
	fmt.Println(anonymousStruct)
}
