package main

import "fmt"

func main() {
	i := 1
	fmt.Println("Write", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	j := 100
	switch {
	case j < 10:
		fmt.Println("less than 10")
	case j > 10:
		fmt.Println("greater than 10")
	}
}
