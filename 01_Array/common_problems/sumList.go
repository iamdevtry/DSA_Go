package main

func sumArray(arr []int) int {
	size := len(arr)
	sum := 0
	for i := 0; i < size; i++ {
		sum += arr[i]
	}

	return sum
}
