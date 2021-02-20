package main

import "fmt"

func main() {
	closeCh := make(chan int, 1)
	go func() {
		for {
			select {
			case value, ok := <-closeCh:
				fmt.Printf("%#v\n", value)
				fmt.Printf("%#v\n", ok)
				if !ok {
					return
				}
			}
		}
	}()

	closeCh <- 1
	close(closeCh)
	for true {

	}
}
