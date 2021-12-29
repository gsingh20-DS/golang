package main

import "fmt"

func factorial(s int) int {
	//l := 2
	if s == 1 || s == 0 {
		return s
	} else {
		return s * factorial(s-1)
	}

}

func main() {
	var i int
	fmt.Print("Enter the number of terms : ")
	fmt.Scan(&i)
	//i := 10
	fmt.Print("factoial is ", factorial(i))

}

//1*2*3
