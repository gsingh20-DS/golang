package main

import (
	"fmt"
)

func myfunc(ch chan int) {
	//

	fmt.Println(234 + <-ch)
	//time.Sleep(2 * time.Second)
}
func main() {
	fmt.Println("start Main method")
	// Creating a channel
	ch := make(chan int, 5)

	println(len(ch))

	go myfunc(ch)
	//ch <- 23
	ch <- 56
	fmt.Println("End Main method")
}
