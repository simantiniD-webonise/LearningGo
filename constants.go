package main

import "fmt"
import "math"

const t string = "I'm constant"

func main() {
	fmt.Println(t)

	const a = 5000000000

	const d = 3e20 / a
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(a))

}
