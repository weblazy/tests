package main

import (
	"fmt"
)

type People struct {
}

type Man struct {
}

func main() {
	a := &People{}
	b := &People{}
	fmt.Printf("%p\n", a)
	fmt.Printf("%p\n", b)
	fmt.Println(a == b)
}
