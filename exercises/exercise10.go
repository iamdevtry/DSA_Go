package main

import "fmt"

// 	Reverse a list in-place. (You cannot use any additional list in other wards
// 	Space Complexity should be O(1). )
// 	Hint: Use two variable, start and end. Start set to 0 and end set to (n-1).
// 	Increment start and decrement end. Swap the values stored at arr[start] and
// 	arr[end]. Stop when start is equal to end or start is greater than end

func Reverse(arr []int) []int {
	size := len(arr)

	start, end := 0, size-1 // 2 pointer
	for start <= end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}

	return arr
}

func main() {
	nums := []int{1, 3, 5, 4, 7, 7, 6, 8, 8, 2}
	nums = Reverse(nums)
	fmt.Println(nums)
}
