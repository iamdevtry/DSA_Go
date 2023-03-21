package main

import "fmt"

func insertAtTheEnd(arr []int, element int) []int {
	return append(arr, element)
}

// Time Complexity: O(1)
// Space Complexity: O(1)
func main() {
	arr := []int{12, 34, 10, 6, 40}
	fmt.Printf("Before added element: %v\n", arr)
	arr = insertAtTheEnd(arr, 99)
	fmt.Printf("After added element: %v\n", arr)
}
