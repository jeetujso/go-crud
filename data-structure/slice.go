package main

import "fmt"

func main() {
	// ::1:: initialize slice
	//x := type{values} COMPOSITE LITERAL (slice is a type of composite literal)
	x := []int{4, 5, 7, 42} // slicec of x which holds integer type value ([]int -> type)
	fmt.Println(x)          // O/P => [4 5 7 42]

	//::2:: Looping of slice

	// len() => retun length of a slice
	fmt.Println(len(x)) // o/p => 4

	//get particular index value
	fmt.Println(x[2]) // o/p => 7
	//for loop ==> use range
	for i, v := range x { // i=> index v=> value and (range over x)
		fmt.Println(i, v) // o/p => 0 4 1 5 2 7 3 42 (index value)
	}

	// slicing of slice
	fmt.Println(x[1:])  // it returns all value from index 1 to end
	fmt.Println(x[1:3]) // it returns all value in between the index value (5,7)

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}

	//::3:: append (Add value to slice)
	x = append(x, 77, 88, 99, 1014) // added value to existing slice x
	fmt.Println(x)                  // [4 5 7 42 77 88 99 1014]

	y := []int{222, 333, 444, 555}

	x = append(x, y...) // appened unlimited number of value (put all values of y to x)
	fmt.Println(x)      // [4 5 7 42 77 88 99 1014 222 333 444 555]
	fmt.Println(len(x)) //12

	// ::4:: deleting slice
	//fmt.Println(x[:2]) // [4 5] up to index 2-1
	//fmt.Println(x[4:]) // index 4 to end
	x = append(x[:2], x[4:]...)
	fmt.Println(x) //[4 5 77 88 99 1014 222 333 444 555]

	//::5:: make()
	z := make([]int, 10, 100) //type, length, capeccity
	fmt.Println(z)            // [0 0 0 0 0 0 0 0 0 0]
	fmt.Println(len(z))       // 10
	fmt.Println(cap(z))       // 100

	z = append(z, 456) //(length and capacity will increase on run time)

	//::6:: multi dimensional slice
	fn := []string{"jeetu", "harpreet", "ankit"}
	fmt.Println(fn)
	ln := []string{"jha", "rupal", "nasir"}
	fmt.Println(ln)

	name := [][]string{fn, ln}
	fmt.Println(name) //[[jeetu harpreet ankit] [jha rupal nasir]]

}
