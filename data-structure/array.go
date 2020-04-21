package main

import "fmt"

func main() {
	//Define emmpty array of lenfth 5 of type ineger
	var x [5]int
	fmt.Println(x) // O/P => [0 0 0 0 0]
	//Assign value to the array
	x[0] = 1
	x[2] = 2
	fmt.Println(x)
	//get length of an array
	fmt.Println(len(x)) // it returns the length of an array

}
