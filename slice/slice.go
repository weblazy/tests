package main

import "fmt"

func main() {
	s := make([]int, 0)
	fmt.Printf("%p\n", s)
	s = append(s, 1)
	fmt.Printf("%p:%p\n", s, &s[0])
	s = append(s, 1)
	fmt.Printf("%p:%p\n", s, &s[0])
	s = append(s, 1)
	fmt.Printf("%p:%p\n", s, &s[0])
	s = append(s, 1)
	fmt.Printf("%p:%p\n", s, &s[0])
}
