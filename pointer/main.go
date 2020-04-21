package main

import "fmt"

func main() {
	// pointer is pointing the some location in memory where the value is stored. address of memory where the value is stored
	a := 45
	fmt.Println(a)   // returns the value :=> 45
	fmt.Println(&a)  // returns the address :=> 0xc000016068
	fmt.Println(*&a) // returns the value of that address :=> 45
}
