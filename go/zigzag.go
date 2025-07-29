// The string "PAYPALISHIRING" is written in a zigzag pattern on a
// given number of rows like this: (you may want to display this pattern in
// a fixed font for better legibility)

// P   A   H   N
// A P L S I I G
// Y   I   R
// And then read line by line: "PAHNAPLSIIGYIR"

// Write the code that will take a string and make this conversion given a number of rows:

// string convert(string s, int numRows);

// Example 1:

// Input: s = "PAYPALISHIRING", numRows = 3
// Output: "PAHNAPLSIIGYIR"
// Example 2:

// Input: s = "PAYPALISHIRING", numRows = 4
// Output: "PINALSIGYAHRPI"
// Explanation:
// P     I    N
// A   L S  I G
// Y A   H R
// P     I
// Example 3:

// Input: s = "A", numRows = 1
// Output: "A"

// Solution

// 📐 The Formula
// Let’s define:

// numRows = total number of rows (e.g., 4)

// r = current row index (starts from 0)

// Then:

// text
// Copy
// Edit
// step1 = 2 * (numRows - r - 1)
// step2 = 2 * r
// These steps are the distances (number of characters to jump) when moving down and up the zigzag.

// You alternate between step1 and step2 in a row unless one of them is 0.

// 📊 Let's Work Through an Example
// We'll use "PAYPALISHIRING" and numRows = 4.

// 🔁 Full Cycle Length
// text
// Copy
// Edit
// cycle = 2 * (numRows - 1) = 2 * 3 = 6
// This cycle repeats every 6 characters.

// 🔹 Row 0 (Top row, r = 0)
// text
// Copy
// Edit
// step1 = 2 * (4 - 0 - 1) = 2 * 3 = 6
// step2 = 2 * 0 = 0
// So for row 0:

// step1 = 6

// step2 = 0

// 🟢 Since step2 = 0, we only use step1 = 6 repeatedly.

// ➡️ Indices: 0, 6, 12...

// 🔹 Row 1 (Second row, r = 1)
// text
// Copy
// Edit
// step1 = 2 * (4 - 1 - 1) = 2 * 2 = 4
// step2 = 2 * 1 = 2
// So:

// First jump = 4

// Next jump = 2

// Then 4, 2, and so on...

// 🟢 You alternate between 4 and 2

// ➡️ Indices: 1, 5, 7, 11, 13...

// 🔹 Row 2 (Third row, r = 2)
// text
// Copy
// Edit
// step1 = 2 * (4 - 2 - 1) = 2 * 1 = 2
// step2 = 2 * 2 = 4
// So:

// First jump = 2

// Next jump = 4

// Then 2, 4, etc.

// 🟢 Again, alternate between 2 and 4

// ➡️ Indices: 2, 4, 8, 10...

// 🔹 Row 3 (Bottom row, r = 3)
// text
// Copy
// Edit
// step1 = 2 * (4 - 3 - 1) = 2 * 0 = 0
// step2 = 2 * 3 = 6
// 🟢 Since step1 = 0, we only use step2 = 6 repeatedly.

// ➡️ Indices: 3, 9, 15...

// 🧮 Final Recap (by row)
// Row	r	step1	step2	Steps Used
// 0	0	6	0	6, 6, 6...
// 1	1	4	2	4, 2, 4, 2...
// 2	2	2	4	2, 4, 2, 4...
// 3	3	0	6	6, 6, 6...

package main

import (
	"fmt"
)

func convertZigzag(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	result := make([]byte, 0, len(s))

	for row := 0; row < numRows; row++ {
		step1 := 2 * (numRows - row - 1)
		step2 := 2 * row
		i := row
		useStep1 := true

		for i < len(s) {
			result = append(result, s[i])

			if step1 == 0 {
				i += step2
			} else if step2 == 0 {
				i += step1
			} else {
				if useStep1 {
					i += step1
				} else {
					i += step2
				}
				useStep1 = !useStep1
			}
		}
	}

	return string(result)
}

func main() {
	s := "PAYPALISHIRING"
	numRows := 4
	fmt.Println(convertZigzag(s, numRows)) // Output: "PINALSIGYAHRPI"
}
