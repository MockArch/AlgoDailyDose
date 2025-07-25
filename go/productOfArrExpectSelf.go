// Problem: Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the
// elements of nums except nums[i].
// You must solve it in O(n) time and without using the division operation.

// input : []int{1, 2, 3, 4}
// output : [24 12 8 6]

package main

import "fmt"

func productOfArrayExpectSelf(arr []int) []int {
	n := len(arr)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		product := 1
		for j := 0; j < n; j++ {
			if j == i {
				continue
			}
			product = product * arr[j]
		}
		res[i] = product
	}
	return res
}

func main() {

	input := []int{1, 2, 3, 4}
	r := productOfArrayExpectSelf(input)
	fmt.Println(r)

}
