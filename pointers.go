package main

import "fmt"

func byVal(ival int) {
	ival = 0
}

func byAddress(ival *int) {
	*ival = 0
}

func main() {
	i := 1
	byVal(i)
	fmt.Println("call by value ", i)
	byAddress(&i)
	fmt.Println("call by reference ", i)
	fmt.Println("address of i ", &i)
}
