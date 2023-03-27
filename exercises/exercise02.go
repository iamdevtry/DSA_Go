package main

import "fmt"

//2) Find the sum of all the elements of a two dimensional list.

func main() {
	arr := [][]int{{1, 2, 3, 4, 6, 3, 4}, {7, 8, 9, 4, 5, 7, 1}}

	sum := 0

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			sum += arr[i][j]
		}
	}

	fmt.Println(sum)
}
