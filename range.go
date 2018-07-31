package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum ", sum)

	var j int
	for i, _ := range nums {
		j = j + i
	}
	fmt.Println("index sum ", j)
}
