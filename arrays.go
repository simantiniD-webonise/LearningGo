package main

import "fmt"

func main() {
	var array [5]int
	fmt.Println("array ", array)

	array[0] = 50
	fmt.Println("new array ", array)
	fmt.Println("element at 0 index ", array[0])

	fmt.Println("lenght ", len(array))

	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(a)
}
