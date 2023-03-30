package main

func power(m, n int) int {
	if n == 0 {
		return 1
	}

	return m * power(m, n-1)
}

func power2(m, n int) int {
	if n == 0 {
		return 1
	} else if n%2 == 0 {
		return power2(m*m, n/2)
	} else {
		return m * power2(m*m, (n-1)/2)
	}
}

func main() {
	println(power(2, 4))
	println(power2(2, 4))
}
