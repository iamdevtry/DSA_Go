package main

func fun(n int) int {
	if n > 100 {
		return n - 10
	}
	return fun(fun(n + 11))
}

func main() {
	r := fun(95)
	println(r)
}
