package main

import "fmt"

func insertElement(a []int, index int, value int) []int {
	n := len(a)
	if index < 0 {
		index = (index%n + n) % n
	}
	switch {
	case index == n: // nil or empty slice or after last element
		return append(a, value)

	case index < n: // index < len(a)
		a = append(a[:index+1], a[index:]...)
		a[index] = value
		return a

	case index < cap(a): // index > len(a)
		a = a[:index+1]
		for i := n; i < index; i++ {
			a[i] = 0
		}
		a[index] = value
		return a

	default:
		b := make([]int, index+1) // malloc
		if n > 0 {
			copy(b, a)
		}
		b[index] = value
		return b
	}
}

func main() {
	// arr := []int{12, 34, 10, 6, 40}
	arr := []int{12}
	// arr := []int{12, 34, 10, 6, 40}
	fmt.Printf("Before added element: %v\n", arr)
	arr = insertElement(arr, 4, 99)
	fmt.Printf("After added element: %v\n", arr)
}
