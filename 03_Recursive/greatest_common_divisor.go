package main

func GCD(m, n int) int {
	if m < n {
		return GCD(n, m)
	}

	if m%n == 0 {
		return n
	}

	return GCD(n, m%n)
}
