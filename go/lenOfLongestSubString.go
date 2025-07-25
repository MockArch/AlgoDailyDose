// Input: "abcabcbb", Length of longest substring without repeating chars: 3
// Input: "bbbbb", Length of longest substring without repeating chars: 1
// Input: "pwwkew", Length of longest substring without repeating chars: 3
// Input: "", Length of longest substring without repeating chars: 0
// Input: "dvdf", Length of longest substring without repeating chars: 3

package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	charIndexMap := make(map[rune]int)
	maxLenght := 0
	left := 0
	for right, char := range s {

		if prevIndex, found := charIndexMap[char]; found && prevIndex >= left {
			left = prevIndex + 1
		}
		charIndexMap[char] = right
		// fmt.Printf("right %v , Char %v   charIndexMap  %v\n", right, char, charIndexMap)

		windowLength := right - left + 1

		if windowLength > maxLenght {

			maxLenght = windowLength

		}

	}
	return maxLenght
}

func main() {

	str := "bbbbb"
	fmt.Println(lengthOfLongestSubstring(str))

}

// Initial State

// s = "abcabcbb"
// charIndexMap = {}
// left = 0, right = 0, maxLength = 0
// Step 0:

// Right: 0 ('a')
// Window: [a]
// Indices: [0]
// charIndexMap: {a:0}
// left: 0, right: 0, maxLength: 1

// Note: Add 'a', window is "a".
// Step 1:

// Right: 1 ('b')
// Window: [a b]
// Indices: [0 1]
// charIndexMap: {a:0, b:1}
// left: 0, right: 1, maxLength: 2

// Note: Add 'b', window is "ab".
// Step 2:

// Right: 2 ('c')
// Window: [a b c]
// Indices: [0 1 2]
// charIndexMap: {a:0, b:1, c:2}
// left: 0, right: 2, maxLength: 3

// Note: Add 'c', window is "abc".
// Step 3:

// Right: 3 ('a') *
// Window: [b c a]
// Indices: [1 2 3]
// charIndexMap: {a:3, b:1, c:2}
// left: 1, right: 3, maxLength: 3

// Note: 'a' seen at 0, move left to 1, window is "bca".
// Legend:

// after char means it's a repeat.
// Window shows the current substring.
// Indices show the positions in the string.
// charIndexMap shows the last seen index for each character.
// Final Output:

// The longest substring without repeating characters is of length 3 (e.g., "abc" or "bca").
