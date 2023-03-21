package main

import "fmt"

//Given an array of integers arr[] of size N and an integer,
//the task is to rotate the array elements to the left by d positions.

// Examples:

// Input:
// arr[] = {1, 2, 3, 4, 5, 6, 7}, d = 2
// Output: 3 4 5 6 7 1 2

// Input: arr[] = {3, 4, 5, 6, 7, 1, 2}, d=2
// Output: 5 6 7 1 2 3 4

func rotate(arr []int, d int) []int {
	return append(arr[d:], arr[:d]...)
}

func main() {
	arr, d := []int{1, 2, 3, 4, 5, 6, 7}, 2
	arr = rotate(arr, d)
	fmt.Println(arr)
}
