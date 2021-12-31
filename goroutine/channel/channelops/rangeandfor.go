package main

import (
	"fmt"
	"sync"
)

func myfunc(ch chan int, wg *sync.WaitGroup) {
	//x := <-ch
	fmt.Println(257 + <-ch)
	wg.Done()
	//close(ch)

}
func main() {
	fmt.Println("start Main method")
	// Creating a channel
	//var ch chan int
	ch := make(chan int, 3)
	wg := sync.WaitGroup{}

	for i := 0; i < cap(ch); i++ {
		//fmt.Println(ch)
		go myfunc(ch, &wg)
		wg.Add(1)
		//wg1.Add(1)
		ch <- i + 32

		//fmt.Println(ch)
	}
	//close(ch)

	println(len(ch))

	//wg1 := sync.WaitGroup{}
	//wg1.Add(len(ch))

	// for res := range ch {
	// 	//println(len(ch))
	// 	//wg1.Add(1)
	// 	fmt.Println(res)
	// 	wg.Done()

	// }

	//
	//
	wg.Wait()
	//wg1.Wait()

	close(ch)
	fmt.Println("End Main method")

}
