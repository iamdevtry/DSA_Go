package main

import "fmt"

// 5) Find the second largest number in the list.
func main() {
	arr := []int{1, 5, 10, 6, 7, 8, 9, 9, 10, 3}

	max := 0
	secondMax := 0
	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			secondMax = max
			max = arr[i]
		} else if arr[i] > secondMax && arr[i] != max {
			secondMax = arr[i]
		}
	}

	fmt.Printf("2nd max %d\n", secondMax)
}
