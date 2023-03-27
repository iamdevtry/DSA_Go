package main

import "fmt"

// 12) Given an array of 0s, 1s and 2s. We need to sort it so that all the 0s are
// before all the 1s and all the 1s are before 2s.
// Hint: Same as above first think 0s and 1s as one group and move all the 2s
// on the right side. Then do a second pass over the array to sort 0s and 1s.
func sort(arr []int) []int {
	size := len(arr)
	l, m, r := 0, 0, size-1 // 3 pointer
	for m <= r {
		if arr[m] == 0 {
			arr[l], arr[m] = arr[m], arr[l]
			l++
			m++
		} else if arr[m] == 1 {
			m++
		} else {
			arr[m], arr[r] = arr[r], arr[m]
			r--
		}
	}

	return arr
}

func main() {
	nums := []int{1, 1, 1, 0, 2, 0, 1, 0, 2, 1}
	nums = sort(nums)
	fmt.Println(nums)
}
