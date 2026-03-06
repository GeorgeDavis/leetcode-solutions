package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	left := 0
	currentSum := 0
	minLen := 1000000 // Or math.MaxInt

	for right := 0; right < len(nums); right++ {
		// 1. Add the current element to the sum
		currentSum += nums[right]

		// 2. While the condition is met (sum >= target),
		// try to SHRINK the window from the left to find a smaller length
		for currentSum >= target {
			// Update minLen here...
			fmt.Println("right:", right, "\n", "left:", left)
			length := right - left + 1
			if length < minLen {
				minLen = length
			}

			// Subtract nums[left] and move left pointer
			// You're subtracting from the sum
			currentSum -= nums[left]
			left++
			fmt.Println("currentSum:", currentSum)
		}
	}
	// Return result (handle case where no subarray was found)
	if minLen == 1000000 {
		return 0
	}
	return minLen
}
