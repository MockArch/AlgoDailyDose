// Given a signed 32-bit integer x, return x with its digits reversed.
// If reversing x causes the value to go outside the signed 32-bit integer
// range [-231, 231 - 1], then return 0.

// Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

// Example 1:

// Input: x = 123
// Output: 321
// Example 2:

// Input: x = -123
// Output: -321
// Example 3:

// Input: x = 120
// Output: 21

//input : 1534236469
// output : 0

// Constraints:

// -231 <= x <= 231 - 1

package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverseInt(x int) int {

	min := math.MinInt32
	max := math.MaxInt32
	isNumberNegative := false

	if x < 0 {
		isNumberNegative = true
		x = -x
	}

	strX := strconv.Itoa(x)
	strX = reverseStr(strX)

	x, _ = strconv.Atoi(strX)
	if isNumberNegative {
		x = -x
	}

	if x >= max || x <= min {
		return 0
	}

	return x
}

func reverseStr(s string) string {
	runeS := []rune(s)

	for i, j := 0, len(runeS)-1; i < j; i, j = i+1, j-1 {
		runeS[i], runeS[j] = runeS[j], runeS[i]
	}
	return string(runeS)
}

func main() {

	inputs := []struct {
		input  int
		output int
	}{
		struct {
			input  int
			output int
		}{
			input:  123,
			output: 321,
		},
		struct {
			input  int
			output int
		}{
			input:  -123,
			output: -321,
		},
		struct {
			input  int
			output int
		}{
			input:  120,
			output: 21,
		},
		struct {
			input  int
			output int
		}{
			input:  1534236469,
			output: 0,
		},
	}

	for _, struc := range inputs {

		result := reverseInt(struc.input)

		if result != struc.output {
			err := fmt.Errorf("Error the expected %d  got output %d  : Input %d", struc.output, result, struc.input)
			panic(err)
		}

	}

}
