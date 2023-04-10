package main

import "fmt"

func fib(n int) int {
	t0 := 0
	t1 := 1
	s := 0

	if n <= 1 {
		return n
	}

	for i := 2; i <= n; i++ {
		s = t0 + t1
		t0 = t1
		t1 = s
	}

	return s
}

func rfib(n int) int {
	if n <= 1 {
		return n
	}

	return rfib(n-2) + rfib(n-1)
}

var F [10]int

func mfib(n int) int {
	if n <= 1 {
		F[n] = n
		return n
	} else {
		if F[n-2] == -1 {
			F[n-2] = mfib(n - 2)
		}
		if F[n-1] == -1 {
			F[n-1] = mfib(n - 1)
		}
		return F[n-2] + F[n-1]
	}
}

func main() {
	fmt.Println(fib(5))
	fmt.Println(rfib(5))

	for i := 0; i < 10; i++ {
		F[i] = -1
	}
	fmt.Println(mfib(5))
}
