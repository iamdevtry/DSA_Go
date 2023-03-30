package main

import "fmt"

var x = 0

func fun1(n int) int {
	if n > 0 {
		x++
		return fun1(n-1) + x
	}

	return 0
}

func main() {
	r := fun1(5)
	fmt.Println(r)
}
