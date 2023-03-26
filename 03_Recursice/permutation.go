package main

import "fmt"

func swap(data []int, x, y int) {
	data[x], data[y] = data[y], data[x]
}

func Permutation(data []int, i int, length int) {
	if length == i {
		fmt.Println(data)
		return
	}

	for j := i; j < length; j++ {
		swap(data, i, j)
		Permutation(data, i+1, length)
		swap(data, i, j)
	}
}

func main() {
	var data [5]int
	for i := 0; i < 5; i++ {
		data[i] = i
	}
	Permutation(data[:], 0, 5)
}
