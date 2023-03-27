package main

// 6) Using AllPermutation function discussed before, write a function, which
// give only distinct solutions.
func distinctPermutations(elements []int) [][]int {
	// Base case: if slice has only one element, return a slice containing it
	if len(elements) == 1 {
		return [][]int{elements}
	}
	// Use a map to keep track of seen elements to avoid duplicates
	seen := make(map[int]bool)
	var permutations [][]int
	for i, elem := range elements {
		if !seen[elem] {
			seen[elem] = true
			// Generate permutations of the remaining elements
			remaining := append([]int{}, elements[:i]...)
			remaining = append(remaining, elements[i+1:]...)
			subPermutations := distinctPermutations(remaining)
			for _, subPermutation := range subPermutations {
				// Append the current element to each sub-permutation
				permutation := append([]int{elem}, subPermutation...)
				// Append the permutation to the result slice
				permutations = append(permutations, permutation)
			}
		}
	}
	return permutations
}
