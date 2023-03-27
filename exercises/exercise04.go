package main

import "fmt"

// 3) Find the largest element in the list.
func main() {
	arr := []int{1, 5, 6, 7, 8, 9, 3}

	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if max > arr[i] {
			max = arr[i]
		}
	}

	fmt.Printf("min %d\n", max)
}
