package main

func SequentialSearch(arr []int, target int) bool {
	size := len(arr)
	for i := 0; i < size; i++ {
		if arr[i] == target {
			return true
		}
	}

	return false
}
