package main

import "fmt"

//Find the largest three distinct elements in an array
//Difficulty Level : Easy

//Given an array with all distinct elements, find the
//largest three elements. Expected time complexity
//is O(n) and extra space is O(1).
// Examples :

// Input: arr[] = {10, 4, 3, 50, 23, 90}
// Output: 90, 50, 23

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]

	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func quickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

func main() {
	arr := []int{12, 13, 1, 10, 34, 1}
	arr = quickSortStart(arr)

	count := 0
	for i := len(arr) - 1; count < 3; i-- {
		if arr[i] != arr[i-1] {
			fmt.Printf("%d ", arr[i])
			count++
		}
	}
	fmt.Print("\n")
}
