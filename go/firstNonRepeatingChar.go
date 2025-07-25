// first non-repeating character in a string

package main

import "fmt"

func firstNonRepeatingChar(s string) rune {

	charCount := map[rune]int{}

	for _, char := range s {
		charCount[char]++
	}

	for _, char := range s {
		if charCount[char] == 1 {
			return char
		}
	}

	return rune(0)
}

func main() {

	str := "leetcode"

	noRepeatativeCharIs := firstNonRepeatingChar(str)
	fmt.Printf("No Repeatative char is %q", noRepeatativeCharIs)
}
