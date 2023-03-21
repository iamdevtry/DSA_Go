package main

import "fmt"

func binarySearch(arr []int, low, high int, value int) int {

	if high < low {
		return -1
	}

	midElement := (low + high) / 2

	if value == arr[midElement] {
		return midElement
	}

	if value > arr[midElement] {
		return binarySearch(arr, (midElement + 1), high, value)
	} else {
		return binarySearch(arr, low, (midElement - 1), value)
	}
}

// Time Complexity: O(log(n)) Using Binary Search
// Auxiliary Space: O(log(n)) due to recursive calls, otherwise iterative version uses Auxiliary Space of O(1).

func main() {
	arr := []int{5, 6, 7, 8, 9, 10}

	position := binarySearch(arr, 0, len(arr)-1, 6)
	if position != -1 {
		fmt.Println("Found!")
		fmt.Printf("Position %d\n", position)
	} else {
		fmt.Println("Not found!")
	}
}
