package main

import "fmt"

func main() {
	m := map[string]int{"k5": 19}

	m["k1"] = 7
	m["k2"] = 15
	fmt.Println("map: ", m)

	v1 := m["k1"]
	fmt.Println("k1", v1)

	_, pr := m["k2"]
	fmt.Println("present ", pr)

}
