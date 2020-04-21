package main

import "fmt"

func main() {
	//fmt.Println("Function provide the modular code")
	foo()
	bar("GOlang")
	w := woo("Reuturing value")
	fmt.Println(w)
	x, y := mouse("chooha", "billy")

	fmt.Println(x, y)
}

// basis syntax :=> func (r receiver) identifire(parameters) (return) { code }

func foo() { //nothing
	fmt.Println("Hello from foo!")
}

func bar(s string) { // take parameters
	fmt.Println("Hello from bar with parameter", s)
}

func woo(s string) string { //take parameter return string
	return fmt.Sprint("Hello from woo with parameter :=>", s)
}

func mouse(fn string, ln string) (string, bool) { //take 2 parameters return 2 values
	a := fmt.Sprint(fn, ln, ", says hello")
	b := false
	return a, b
}
