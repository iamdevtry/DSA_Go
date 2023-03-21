package main

import "fmt"

// delete an element from a sorted array
func delete(arr []int, n int) []int {
	//find the index of the element to be deleted
	i := search(arr, n)

	//if the element is not present in the array
	if i == -1 {
		return arr
	}

	//shift all the elements to the left
	for j := i; j < len(arr)-1; j++ {
		arr[j] = arr[j+1]
	}

	//return the array
	return arr[:len(arr)-1]
}

// search an element in a sorted array
func search(arr []int, n int) int {
	//start from the first element
	i := 0

	//loop until the end of the array
	for i < len(arr) {
		//if the element is found
		if arr[i] == n {
			//return the index
			return i
		}

		//increment the index
		i++
	}

	//if the element is not found
	return -1
}

// main function
func main() {
	//declare and initialize the array
	arr := []int{10, 20, 30, 40, 50}

	//print the array
	fmt.Println(arr)

	//delete the element
	arr = delete(arr, 30)

	//print the array
	fmt.Println(arr)
}
