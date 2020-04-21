package main

import "fmt"

func main() {
	//variadic parameters in function
	foot(1, 2, 3, 4, 5, 6, 7, 8)

	sum := add(4, 5, 6, 7, 8)
	fmt.Println(sum)

	xi := []int{5, 4, 3, 2, 1}
	sum2 := add(xi...)
	fmt.Println("sum2", sum2)
}

func foot(x ...int) { // it takes unlimited nunber of int parameters
	fmt.Println(x)        //[1 2 3 4 5 6 7 8]
	fmt.Printf("%T\n", x) //[]int :=> slice of int
}
func add(y ...int) int {
	sum := 0
	for i, v := range y {
		sum += v
		fmt.Println("value id ->", v, "and sum is =>", sum, i)
	}
	return sum
}
