package main

import "fmt"

func insertionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i; j >= 1 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}

	return arr
}

func main() {
	var n = []int{1, 39, 2, 9, 7, 54, 11}

	fmt.Println(insertionSort(n))
}
