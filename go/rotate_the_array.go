// Given an array, rotate the array to the right by k steps, where k is non-negative.
// Input = []int{1, 2, 3, 4, 5, 6, 7}
// k= 3
// output = [5 6 7 1 2 3 4]

package main

import "fmt"

func reverse(arr []int, start, end int) {

	for start < end {

		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}

func rotateArray(arr []int, k int) {

	lenArr := len(arr)

	if lenArr == 0 {
		return
	}

	k %= lenArr

	reverse(arr, 0, lenArr-1)
	reverse(arr, 0, k-1)
	reverse(arr, k, lenArr-1)
}

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7}
	// [5 6 7 1 2 3 4]
	fmt.Println("Hello ! Rotate the array input ", input)
	rotateArray(input, 3)
	fmt.Println("after rotation", input)
}
