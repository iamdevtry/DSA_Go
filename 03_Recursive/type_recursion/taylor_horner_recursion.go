package main

import "fmt"

var r float32

func e(x, n int) float32 {
	if n == 0 {
		return r
	}

	r = 1 + float32(x)*r/float32(n)
	return e(x, n-1)
}

func main() {
	fmt.Printf("%f \n", e(2, 10))
}
