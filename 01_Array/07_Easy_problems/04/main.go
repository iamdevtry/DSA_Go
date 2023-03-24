package main

import "fmt"

// Rearrange array such that even positioned are greater than odd
// Difficulty Level : Easy
// Given an array A of n elements, sort the array according to the following relations :

// A[i] >= A[i-1]                     , if i is even,  ∀ 1 <= i < n
// A[i] <= A[i-1]                     , if i is odd ,  ∀ 1 <= i < n
// Print the resultant array.

// Examples :

// Input : A[] = {1, 2, 2, 1}
// Output :  1 2 1 2
// Explanation :
// For 1st element, 1  1, i = 2 is even.
// 3rd element, 1  1, i = 4 is even.

// Input : A[] = {1, 3, 2}
// Output : 1 3 2
// Explanation :
// Here, the array is also sorted as per the conditions.
// 1  1 and 2 < 3.

func main() {
	arr := []int{1, 2, 2, 1}
	for i := 1; i < len(arr); i++ {
		if i%2 == 0 {
			if arr[i] > arr[i-1] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
			}
		} else {
			if arr[i] < arr[i-1] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
			}
		}
	}

	fmt.Println(arr)
}
