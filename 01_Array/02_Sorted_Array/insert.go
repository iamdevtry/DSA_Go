package main

import (
	"fmt"
	"sort"
)

func main() {
	data := []int{20, 21, 22, 24, 25, 26, 28, 29, 30, 31, 32}
	var items = []int{23, 27}
	for _, x := range items {
		i := sort.Search(len(data), func(i int) bool { return data[i] >= x })
		if i < len(data) && data[i] == x {
			fmt.Println(i)
		} else {
			data = append(data, 0)
			copy(data[i+1:], data[i:])
			data[i] = x
		}
		fmt.Println(data)
	}
}

// Time Complexity: O(N) [In the worst case all elements may have to be moved]
// Auxiliary Space: O(1)
