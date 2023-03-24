package main

import "fmt"

// Find Second largest element in an array
// Difficulty Level : Easy

// Given an array of integers, our task is to write a program that efficiently finds the second-largest element present in the array.

// Example:

// Input: arr[] = {12, 35, 1, 10, 34, 1}
// Output: The second largest element is 34.
// Explanation: The largest element of the
// array is 35 and the second
// largest element is 34

// Input: arr[] = {10, 5, 10}
// Output: The second largest element is 5.
// Explanation: The largest element of
// the array is 10 and the second
// largest element is 5

// Input: arr[] = {10, 10, 10}
// Output: The second largest does not exist.
// Explanation: Largest element of the array
// is 10 there is no second largest element

// Solved this problem with bubble sort
func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		swapped := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return arr
}

func main() {
	arr := []int{12, 13, 1, 10, 34, 1}
	arr = bubbleSort(arr)

	i := len(arr) - 1
	for i >= 0 {
		if arr[i] != arr[i-1] {
			fmt.Printf("%d ", arr[i-1])
			break
		}
		i--
	}
	fmt.Print("\n")
}
