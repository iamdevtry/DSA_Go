package main

import "fmt"

func findElement(arr []int, key int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == key {
			return i
		}
	}
	return -1
}

func deleteElement(arr []int, element int) []int {
	pos := findElement(arr, element)
	if pos == -1 {
		fmt.Println("Element not found")
		return arr
	}

	// Deleting element
	return append(arr[:pos], arr[pos+1:]...)
}

// Time Complexity: O(N)
// Space Complexity: O(1)

func main() {
	arr := []int{12, 34, 10, 6, 40}
	fmt.Printf("Before deleted element: %v\n", arr)
	arr = deleteElement(arr, 10)
	fmt.Printf("After deleted element: %v\n", arr)
}
