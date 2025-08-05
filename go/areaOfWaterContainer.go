// Container With Most Water

// Hint
// You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

// Find two lines that together with the x-axis form a container, such that the container contains the most water.

// Return the maximum amount of water a container can store.

// Notice that you may not slant the container.

// Example 1:

// Input: height = [1,8,6,2,5,4,8,3,7]
// Output: 49
// Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
// Example 2:

// Input: height = [1,1]
// Output: 1

// Constraints:

// n == height.length
// 2 <= n <= 105
// 0 <= height[i] <= 104

// Solution

// The formula min(height[i], height[j]) * (j - i) calculates the area of water a container can hold between two vertical lines at positions i and j.

// Why?
// min(height[i], height[j]) — because the shorter line limits how tall the container can be (water spills over the shorter side).

// (j - i) — because this is the width between the two lines (how far apart they are on the x-axis).

// Example:
// For height = [1, 8, 6, 2, 5, 4, 8, 3, 7]:

// Taking i = 0, j = 8:

// height[0] = 1, height[8] = 7

// So, height = min(1, 7) = 1

// width = 8 - 0 = 8

// Area = 1 * 8 = 8

// This gives the water area between those two lines.

// Let me know if you want a visual graph generated for this!

package main

import (
	"fmt"
)

func maxArea(height []int) int {
	l := len(height)
	maxAmountofWater := 0
	for left, right := 0, l-1; right > left; {
		Currentheight := min(height[left], height[right])
		Currentwidth := right - left
		Currentarea := Currentheight * Currentwidth

		if Currentarea > maxAmountofWater {
			maxAmountofWater = Currentarea
		}

		if height[left] < height[right] {
			left = left + 1
		} else {
			right = right - 1
		}
	}
	return maxAmountofWater
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	testCases := [][]int{
		{1, 8, 6, 2, 5, 4, 8, 3, 7}, // Expected: 49
		{1, 1},                      // Expected: 1
		{4, 3, 2, 1, 4},             // Expected: 16
		{1, 2, 1},                   // Expected: 2
		{1, 3, 2, 5, 25, 24, 5},     // Expected: 24
	}

	for i, tc := range testCases {
		fmt.Printf("Test case %d: %v\n", i+1, tc)
		fmt.Printf("Max area = %d\n", maxArea(tc))
		fmt.Println("----------------------------")
	}
}
