package main

import "fmt"

// Given a string s, find the length of the longest substring without duplicate characters.

func longstring(s string) int {
	m := make(map[byte]int)
	high := 0
	left := 0

	for right := 0; right < len(s); right++ {
		if val, exists := m[s[right]]; exists {
			left = max(left, val+1)
		}
		m[s[right]] = right
		high = max(high, right-left+1)
	}
	return high
}

func max(a, b int) int {

	if a > b {
		return a
	}
	return b
}

func main() {
	s := ""
	fmt.Println(longstring(s))
}

// Example 1:

// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
// Example 2:

// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
