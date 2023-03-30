package main

func fun(n int) {
	if n > 0 {
		println(n)
		fun(n - 1)
		fun(n - 1)
	}
}

func main() {
	fun(3)
}
