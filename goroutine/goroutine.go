package main

import (
	"fmt"
	"time"
)

func display(str string) {
	for w := 0; w < 1; w++ {
		abc := time.Now().Format("2013-02-31")
		time.Sleep(100 * time.Microsecond)
		fmt.Printf("\ncurrent time is %s", (abc))
		fmt.Printf("\nvalue %s", (str))
	}
}

func main() {

	// Calling Goroutine
	go display("Welcome")
	go func() {
		fmt.Printf("This is anonymous go routine")
	}()
	time.Sleep(100 * time.Microsecond)
	// Calling normal function
	display("GeeksforGeeks")
}
