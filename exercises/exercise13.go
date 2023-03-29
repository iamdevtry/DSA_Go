package main

import "fmt"

// Find the duplicate elements in a list of size n where each element is in the
// range 0 to n-1.
// Hint:
// Approach 1: Compare each element with all the elements of the list (using
// two loops) O(n2) solution

func findDuplicate2Loop(arr []int) {
	size := len(arr)
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if arr[i] == arr[j] {
				fmt.Printf("duplicate %d\n", arr[i])
			}
		}
	}
}

// Approach 2: Maintain a Hash-Table. Set the hash value to 1 if we encounter
// the element for the first time. When we same value again we can see that the
// hash value is already 1 so we can print that value. O(n) solution, but
// additional space is required.

func findDuplicateHashTable(arr []int) {
	exist := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		exist[arr[i]] += 1
	}

	for k, v := range exist {
		if v > 1 {
			fmt.Printf("%d duplicates %d times\n", k, v)
		}
	}

}

// Approach 3: We will exploit the constraint "every element is in the range 0
// to n-1". We can take a list arr[] of size n and set all the elements to 0.
// Whenever we get a value say val1. We will increment the value at arr[var1]
// index by 1. In the end, we can traverse the list arr and print the repeated
// values. Additional Space Complexity will be O(n) which will be less than
// Hash-Table approach.
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 7, 7, 8, 9, 0}
	// findDuplicate1(arr)
	findDuplicateHashTable(arr)
}
