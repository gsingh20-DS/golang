//func make([]T, len, cap) []T

package main

import "fmt"

func main() {
	var slice = make([]int, 5, 8)

	//slice := []int{1, 2, 3, 4, 5}

	for i, j := range slice {
		println(i, slice[j])
	}
	fmt.Printf("\nthe capacity is %d", cap(slice))
	fmt.Printf("\nthe length is %d", len(slice))
}
