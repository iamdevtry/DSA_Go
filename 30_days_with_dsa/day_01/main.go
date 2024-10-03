package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	// init an array of 5 elements
	arr := [5]int{1, 2, 3, 4, 5}

	fmt.Println("Array:", arr)
	fmt.Println("Length of array: ", len(arr))

	// traverse the array
	for i := 0; i < len(arr); i++ {
		fmt.Println("Element at index", i, "is", arr[i])
	}

	// find an element in the array
	for i := 0; i < len(arr); i++ {
		if arr[i] == 3 {
			fmt.Println("Element 3 found at index", i)
		}
	}
}
