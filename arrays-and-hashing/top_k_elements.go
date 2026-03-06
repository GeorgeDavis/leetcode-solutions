package main

import "sort"

func topKFrequent(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return nil
	}

	numsCounter := make(map[int]int, len(nums)) // bounded the map memory allocation

	// must iterate over nums and count the frequency of each element
	// for loop is O(n) time complexity
	for _, num := range nums {
		numsCounter[num]++
	}

	// flatten into a slice
	numsSlice := make([]int, 0, len(numsCounter))
	for key := range numsCounter {
		numsSlice = append(numsSlice, key)
	}

	sort.Slice(numsSlice, func(i, j int) bool {
		return numsCounter[numsSlice[i]] > numsCounter[numsSlice[j]]
	})

	return numsSlice[:k]
}
