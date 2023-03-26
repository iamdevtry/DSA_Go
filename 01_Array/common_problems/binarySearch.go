package main

import "sort"

// binary search in sorted list
func binarySearch(arr []int, target int) bool {
	size := len(arr)
	sort.Ints(arr)

	low, high := 0, size-1

	var mid int
	for low <= high {
		mid = low + (high-low)/2
		if arr[mid] == target {
			return true
		} else {
			if arr[mid] < target {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	return false
}
