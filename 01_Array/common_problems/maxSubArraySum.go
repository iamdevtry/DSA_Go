package main

import "fmt"

func maxSubArraySum(arr []int) int {
	size := len(arr)
	result := 0
	cur := 0

	for i := 0; i < size; i++ {
		cur += arr[i]
		if cur < 0 {
			cur = 0
		}

		if result < cur {
			result = cur
		}
	}

	return result
}

func main() {
	data := []int{1, -2, 3, 4, -4, 6, -14, 8, 2}
	fmt.Println("Max sub array sum :", maxSubArraySum(data))
}
