package main

import "fmt"

func printSubArrays(arr []int, start int, end int) {
	if end == len(arr) {
		return
	} else if start > end {
		printSubArrays(arr, 0, end+1)
	} else {
		fmt.Print("[")
		for i := start; i < end; i++ {
			fmt.Printf("%d, ", arr[i])
		}

		fmt.Printf("%d]\n", arr[end])

		printSubArrays(arr, start+1, end)
	}
}

func main() {
	arr := []int{1, 2, 3}
	printSubArrays(arr, 0, 0)
}
