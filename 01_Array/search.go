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

// Time Complexity: O(N)
// Space Complexity: O(1)
func main() {
	arr := []int{12, 34, 10, 6, 40}

	position := findElement(arr, 8)
	if position != -1 {
		fmt.Println("Found!")
		fmt.Printf("Position %d\n", position)
	} else {
		fmt.Println("Not found!")
	}
}
