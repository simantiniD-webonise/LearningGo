package main

import "fmt"

func main() {
	var s = make([]string, 3)
	fmt.Println(s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("slice s is ", s)

	fmt.Println("slice length ", len(s))

	s = append(s, "d")
	s = append(s, "e")
	s = append(s, "f", "g", "h", "i", "j")
	fmt.Println("appended string ", s)

	l := s[2:5]
	fmt.Println("l", l)

	b := s[:5]
	fmt.Println("b", b)

}
