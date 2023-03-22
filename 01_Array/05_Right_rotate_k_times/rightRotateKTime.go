package main

import "fmt"

func rightRotate(arr []int, k int) []int {
	for i := 0; i < k; i++ {
		arr = append(arr[len(arr)-1:], arr[:len(arr)-1]...)
	}
	return arr
}

//Time Complexity: O(n*k)
//Auxiliary Space: O(1)

func RightRotate(arr []int, n, k int) []int {

	k = k % n

	for i := 0; i < n; i++ {
		if i < k {

			arr[i] = arr[n+i-k]
		} else {

			arr[i] = arr[i-k]
		}
	}
	return arr
}

//Time Complexity: O(n)
//Auxiliary Space: O(1)

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{1, 2, 3, 4, 5}
	k := 4

	arr1 = RightRotate(arr1, len(arr1), k)
	arr2 = rightRotate(arr2, k)
	fmt.Println(arr1)
	fmt.Println(arr2)
}
