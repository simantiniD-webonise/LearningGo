package main

import "fmt"

/**
-structs
-methods
**/
type person struct {
	name string
	age  int
}

type rect struct {
	height int
	width  int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	fmt.Println(person{"Sheldon", 30})
	fmt.Println(person{"Leonard", 29})
	fmt.Println(person{"Penny", 28})

	s := person{"Howard", 40}
	fmt.Println("Howard name", s.name)

	sp := &s
	fmt.Println("Howard age before", sp.age)

	sp.age = 50

	fmt.Println("Howard age after", sp.age)

	r := rect{width: 10, height: 5}
	fmt.Println("Area ", r.area())
	fmt.Println("Perim ", r.perim())

	rp := &r

	fmt.Println("Area ", rp.area())
	fmt.Println("Perim ", rp.perim())
}
