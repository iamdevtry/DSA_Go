package main

import "fmt"

func SortZeroAndOne(arr []int) []int {
	size := len(arr)
	start, end := 0, size-1

	for start <= end {
		if arr[start] == 0 {
			start++
		} else if arr[end] == 1 {
			end--
		} else {
			arr[start], arr[end] = arr[end], arr[start]
			start++
			end--
		}
	}

	return arr
}

func main() {
	nums := []int{1, 1, 1, 0, 0, 0, 1, 0, 0, 1}
	nums = SortZeroAndOne(nums)
	fmt.Println(nums)
}
