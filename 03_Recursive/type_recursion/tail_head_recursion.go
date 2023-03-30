package main

// Head Recursion
func fun1(n int) {
	if n > 0 {
		fun1(n - 1)
		println(n)
	}
}

// Tail Recursion
func fun2(n int) {
	if n > 0 {
		println(n)
		fun2(n - 1)
	}
}

func main() {
	fun1(3)
	fun2(3)
}
