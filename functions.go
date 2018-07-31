package main

import "fmt"

/**
-functions
-multiple return values
-variadic functions
-closures (anonymous functions)
**/
func main() {
	a := sumAB(10, 20)
	fmt.Println("sum is ", a)

	b := sum(5, 10, 5)
	fmt.Println("sum 2 is ", b)

	c, d := multValues()
	fmt.Println("c ", c)
	fmt.Println("d ", d)

	_, e := multValues()
	fmt.Println("1st val", e)

	f, _ := multValues()
	fmt.Println("2nd val", f)

	sumVariadicFn(1, 2)
	sumVariadicFn(1, 2, 3)

	g := []int{1, 2, 3, 4, 5}
	sumVariadicFn(g...)

	nextInt := initSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())

}

func sumAB(a int, b int) int {
	return a + b
}

func sum(a, b, c int) int {
	return a + b + c
}

func multValues() (int, int) {
	return 1, 2
}

func sumVariadicFn(nums ...int) {
	fmt.Println("nums ", nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("total ", total)
}

func initSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
