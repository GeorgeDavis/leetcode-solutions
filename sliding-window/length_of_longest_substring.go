package main

import "fmt"


func lengthOfLongestSubstring(s string) int {
	// Map to store the frequency of characters in the current window
	m := make(map[byte]int)
	left := 0
	maxLen := 0

	for right := 0; right < len(s); right++ {
		// string indexing uses the same syntax as maps
		char := s[right]
		fmt.Println("current right:", char)
		m[char]++

		// If we have a duplicate, shrink the window from the left
		for m[char] > 1 {
			leftChar := s[left]
			m[leftChar]--
			left++
		}

		// Calculate current window size
		currentWindowSize := right - left + 1
		if currentWindowSize > maxLen {
			maxLen = currentWindowSize
		}
	}

	return maxLen
}
