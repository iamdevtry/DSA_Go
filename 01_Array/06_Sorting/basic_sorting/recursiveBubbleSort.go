package main

import "fmt"

func recursiveBubbleSort(arr []int, size int) []int {
	if size == 1 {
		return arr
	}

	for i := 0; i < size-1; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}

	recursiveBubbleSort(arr, size-1)
	return arr
}

func main() {
	var n = []int{1, 39, 2, 9, 7, 54, 11}

	fmt.Println(recursiveBubbleSort(n, len(n)))
}
