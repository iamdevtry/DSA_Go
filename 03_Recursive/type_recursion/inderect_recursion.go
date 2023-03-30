package main

func funA(n int) {
	if n > 0 {
		println(n)
		funB(n - 1)
	}
}

func funB(n int) {
	if n > 1 {
		println(n)
		funA(n / 2)
	}
}

func main() {
	funA(20)
}
