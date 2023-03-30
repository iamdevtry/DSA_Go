package main

func factory(n int) int {
	if n == 0 {
		return 1
	}

	return factory(n-1) * n
}

func main() {
	println(factory(5))

}
