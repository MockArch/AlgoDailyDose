// Median of Two Sorted Arrays

// Given two sorted arrays nums1 and nums2 of size m and n respectively,
// return the median of the two sorted arrays.

// The overall run time complexity should be O(log (m+n)).

// Example 1:

// Input: nums1 = [1,3], nums2 = [2]
// Output: 2.00000
// Explanation: merged array = [1,2,3] and median is 2.
// Example 2:

// Input: nums1 = [1,2], nums2 = [3,4]
// Output: 2.50000
// Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.

// Constraints:

// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -106 <= nums1[i], nums2[i] <= 106

package main

import (
	"fmt"
	"math"
)

func QuickSort(arr []int, low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)
		QuickSort(arr, low, pivotIndex-1)
		QuickSort(arr, pivotIndex+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
}

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	merged := append(nums1, nums2...)
	lenMerged := len(merged)
	QuickSort(merged, 0, lenMerged-1)
	mid := int(math.Ceil(float64(lenMerged) / 2))

	if mid%2 != 0 {
		return float64(merged[mid-1])
	}
	return float64(merged[mid]+merged[mid-1]) / 2
}

func main() {
	testingKit := []struct {
		nums1    []int
		nums2    []int
		expected float64
	}{
		{
			nums1:    []int{1, 3},
			nums2:    []int{2},
			expected: 2.00000,
		},
		{
			nums1:    []int{1, 2},
			nums2:    []int{3, 4},
			expected: 2.50000,
		},
		{
			nums1:    []int{0, 0, 0, 0, 0},
			nums2:    []int{-1, 0, 0, 0, 0, 0, 1},
			expected: 0,
		},
	}

	for _, strukt := range testingKit {
		result := findMedianSortedArrays(strukt.nums1, strukt.nums2)
		fmt.Printf("nums1 : %v  nums : %v  median is %f   Expected %f \n", strukt.nums1, strukt.nums2, result, strukt.expected)

		if result != strukt.expected {
			fmt.Errorf("Expected %f  and we got result is %f", strukt.expected, result)
		}
	}
}
