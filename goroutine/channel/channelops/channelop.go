// package main

// import (
// 	"fmt"
// 	"time"
// )

// func myfunc(ch chan int, ch1 chan int) {
// 	time.Sleep(1 * time.Second)
// 	fmt.Println(257 + <-ch)
// 	time.Sleep(1 * time.Second)
// 	fmt.Println(267 + <-ch1)
// }
// func main() {
// 	fmt.Println("start Main method")
// 	// Creating a channel
// 	var ch1 chan int
// 	ch := make(chan int)

// 	go myfunc(ch, ch1)
// 	ch1 <- 56
// 	ch <- 23

// 	fmt.Println("End Main method")
// }

package main

import (
	"fmt"
	"time"
)

func myfunc(ch chan int) {

	fmt.Println(257 + <-ch)
	time.Sleep(2 * time.Second)
	//close(ch)
}
func main() {
	fmt.Println("start Main method")
	// Creating a channel
	//var ch chan int
	ch := make(chan int, 3)

	go myfunc(ch)
	// for i := range ch {
	// 	fmt.Println(i)
	ch <- 2
	time.Sleep(2 * time.Second)

	ch <- 4
	time.Sleep(2 * time.Second)
	close(ch)
	// }
	fmt.Println(cap(ch))
	fmt.Println("End Main method")

}
