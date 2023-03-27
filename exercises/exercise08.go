package main

// 	Print all the maxima’s in a list. (A value is a maximum if the value before
// 	and after its index are smaller than it is or does not exist.)
// 	Hint: Start traversing list from the end and keep track of the max element.
// 	If we encounter an element whose value is grater then max, print the element
// 	and update max.

import "fmt"

func printMaximas(nums []int) {
	n := len(nums)
	max := nums[n-1]

	for i := n - 2; i >= 0; i-- {
		if nums[i] > max {
			fmt.Println(nums[i])
			max = nums[i]
		}
	}
}

func main() {
	nums := []int{1, 3, 5, 4, 7, 7, 6, 8, 8, 2}
	printMaximas(nums)
}
