package main

import "fmt"

func RotateArray(arr []int, k int) {

}

func ReverseArray(arr []int, start int, end int) {
	i, j := start, end
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func main() {
	arr := []int{1, 5, 6, 7, 2, 3}
	ReverseArray(arr, 2, len(arr)-1)
}
