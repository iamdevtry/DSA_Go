package main

import (
	"fmt"
)

func e(x, n int) float64 {
	var s float64 = 1
	var num float64 = 1
	var den float64 = 1
	for i := 1; i <= n; i++ {
		num *= float64(x)
		den *= float64(i)
		s += float64(num / den)
	}
	return s
}

func main() {
	fmt.Printf("%f \n", e(2, 10))
}
