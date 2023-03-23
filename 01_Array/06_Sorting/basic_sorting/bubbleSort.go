package main

import "fmt"

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j] = arr[j+1]
			}
		}
	}

	return arr
}

func bubbleSortWithPseudo(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		swapped := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}

	return arr
}

func main() {
	arr := []int{1, 39, 2, 9, 7, 54, 11}

	arr = bubbleSort(arr)

	fmt.Println(arr)
}
