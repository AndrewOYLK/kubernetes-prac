package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}

	// copy(dst, src)

	// copy(slice1, slice2)
	// fmt.Println(slice1, slice2) // [5,4,3,4,5] [5,4,3]

	copy(slice2, slice1)
	fmt.Println(slice1, slice2) // [1,2,3,4,5] [1,2,3]
}
