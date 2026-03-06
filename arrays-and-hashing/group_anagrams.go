package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, word := range strs {
		chars := []byte(word)
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j] // using this sort algo we can sort each word into the same key each time
		})

		// get chars back into string
		key := string(chars)
		groups[key] = append(groups[key], word)
	}

	result := [][]string{}
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}
