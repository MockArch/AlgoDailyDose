// Package main provides a function to merge two sorted integer slices into a single sorted slice.

// The core function, merge(a, b []int) []int, takes two sorted slices of integers and returns
// a new slice containing all elements from both inputs in sorted order. This is commonly used in
// merge sort and other algorithms requiring sorted data combination.

// Algorithm:
// - Uses two pointers to traverse both input slices.
// - Compares elements at each pointer, appending the smaller to the output.
// - Appends any remaining elements from either slice after one is exhausted.

// Time Complexity: O(n + m), where n and m are the lengths of the input slices.

// Edge Cases:
// - Handles empty slices.
// - Works for slices of different lengths.

// Example:
//     a := []int{1, 3, 6}
//     b := []int{2, 4, 5}
//     merged := merge(a, b) // merged = [1 2 3 4 5 6]

package main

import "fmt"

func merge(a, b []int) []int {

	out := make([]int, 0, len(a)+len(b))
	i := 0
	j := 0

	for i < len(a) && j < len(b) {

		if a[i] < b[j] {
			out = append(out, a[i])
			i++
		} else {
			out = append(out, b[j])
			j++
		}
	}

	out = append(out, a[i:]...)
	out = append(out, b[j:]...)

	return out

}

func main() {

	a := []int{1, 3, 6, 7, 10}
	b := []int{0, 2, 5, 8, 9, 11, 15}

	mergedSlice := merge(a, b)

	fmt.Println("Merged Slices is ", mergedSlice)
}
