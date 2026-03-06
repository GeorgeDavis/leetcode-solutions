package main

import "strings"

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}

	var filtered []byte
	for i := 0; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= '0' && s[i] <= '9') {
			filtered = append(filtered, []byte(strings.ToLower(string(s[i])))[0])
		}
	}

	charSlice := filtered
	for i, j := 0, len(charSlice)-1; i < j; i, j = i+1, j-1 {
		if charSlice[i] != charSlice[j] {
			return false
		}
	}
	return true
}
