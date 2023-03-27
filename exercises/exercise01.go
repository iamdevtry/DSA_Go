package main

import "fmt"

//1) Find average of all the elements in a list.

func main() {
	arr := []int{3, 5, 6, 7, 8, 2, 3}

	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	fmt.Println(sum / len(arr))
}
