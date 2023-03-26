package main

import "fmt"

func factorial(i int) int {
	//Termination Condition
	if i <= 1 {
		return 1
	}

	//Body, Recursive Expansion
	return i * factorial(i-1)
}

func main() {
	fmt.Println("factorial 5 is :: ", factorial(5))
}
