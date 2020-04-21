package main

import "fmt"

func main() {
	fmt.Println("map")
	// map is a key value pair
	//mobile := map[string]int{values} // composite literal with type map[string]int
	m := map[string]int{
		"jeetu":    28,
		"harpreet": 30,
		"ankit":    27,
	}

	fmt.Println(m) // map[ankit:27 harpreet:30 jeetu:28]

	// access map with key
	fmt.Println(m["jeetu"]) //28

	fmt.Println(m["charan"]) // it returns value 0 when key not found

	//checking key
	v, ok := m["charan"]
	fmt.Println(v)  // 0
	fmt.Println(ok) // false

	if v, ok := m["charan"]; ok {
		fmt.Println("This is if conditin", v)
	}

	//::Add map key value
	m["baljeet"] = 33
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}

	//::Delete map key value
	delete(m, "baljeet")
	fmt.Println(m)

	delete(m, "abc") // no error
	fmt.Println(m)
	if v, ok := m["ankit"]; ok {
		fmt.Println("deleted", v)
		delete(m, "ankit")
		fmt.Println(m)
	}
}
