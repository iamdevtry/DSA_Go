package main

import "fmt"

var (
	p = 1
	f = 1
)

func e(x, n int) float32 {
	var r float32
	if n == 0 {
		return 1
	}

	r = e(x, n-1)
	p = p * x
	f = f * n
	return r + float32(p)/float32(f)
}

func main() {
	fmt.Printf("%f \n", e(3, 10))
}
