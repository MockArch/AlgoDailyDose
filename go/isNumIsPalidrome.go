// Palindrome Number

// Given an integer x, return true if x is a palindrome, and false otherwise.

// Example 1:

// Input: x = 121
// Output: true
// Explanation: 121 reads as 121 from left to right and from right to left.
// Example 2:

// Input: x = -121
// Output: false
// Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
// Example 3:

// Input: x = 10
// Output: false
// Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

// Constraints:

// -231 <= x <= 231 - 1

package main

import (
	"fmt"
	"math"
)

func reverseInteger(num int) int {
	// Handle the sign
	sign := 1
	if num < 0 {
		sign = -1
		num = -num
	}

	reversed := 0
	for num > 0 {
		// Extract the last digit
		digit := num % 10
		// Shift the reversed number to the left and add the digit
		reversed = reversed*10 + digit
		// Remove the last digit from the original number
		num /= 10
	}

	// Apply the sign to the reversed number
	reversed *= sign

	// Handle overflow/underflow
	if reversed > math.MaxInt32 || reversed < math.MinInt32 {
		return 0
	}

	return reversed
}

func main() {
	testCases := []int{
		123,
		-123,
		120,
		0,
		1534236469,
	}

	for _, testCase := range testCases {
		fmt.Printf("reverseInteger(%d) = %d\n", testCase, reverseInteger(testCase))
	}
}
