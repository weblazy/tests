package main

import (
	"fmt"
	"time"
)

type st struct {
	field string
}

func main() {
	a := []st{st{field: "1"}, st{field: "2"}}
	a[0].field = "3"
	fmt.Printf("%#v\n", a)
	defer fmt.Printf("%#v\n", "hahhh")
	time.Sleep(1 * time.Minute)
}
