package main

import "fmt"

func sumOfDigit(s int) int {
	//l := 2
	res := 0
	for s > 0 {
		reminder := s % 10
		res = res + reminder
		s /= 10

	}
	return (res)
}

func main() {
	var i int
	fmt.Print("\nEnter the number of terms : ")
	fmt.Scan(&i)
	//i := 10

	fmt.Print("\n sum of digit is ", sumOfDigit(i))

}

//1234 === 4321
