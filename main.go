package main

import "fmt"

func main() {
	// input := "pwwkew"
	// result := lengthOfLongestSubstring(input)
	// fmt.Printf("The length of the longest substring for '%s' is: %d\n", input, result)

	prices1 := []int{7, 1, 5, 3, 6, 4}
	// prices2 := []int{7, 6, 4, 3, 1}
	result := maxProfit(prices1)
	fmt.Println("result:", result)

	// target := 7
	// nums := []int{1, 1, 1, 1, 7}
	// result := minSubArrayLen(target, nums)
	// fmt.Println("result:", result)

}

func maxProfit(prices []int) int {
	maxProfit := 0

	// brute force technique
	lowestPrice := 100
	potentialProfit := 0
	for i := 0; i < len(prices); i++ {
		currentPrice := prices[i]
		if currentPrice < lowestPrice {
			lowestPrice = currentPrice
			fmt.Println("lowestPrice:", lowestPrice)
		}

		// calculate potentialProfit = currentPrice - lowestPrice
		// if potentialProfit > maxProfit set maxProfit == potentialProfit
		potentialProfit = currentPrice - lowestPrice
		if potentialProfit > maxProfit {
			maxProfit = potentialProfit
			fmt.Println("maxProfit:", maxProfit)
		}
	}


	// sliding window technique
	left := 0
	right := 1
	for right < len(prices) {
		if prices[right] > prices[left] {
			profit := prices[right] - prices[left]
			if profit > maxProfit {
				maxProfit = profit
			}
		} else {
			left = right
		}
		right++
	}

	fmt.Println("maxProfit:", maxProfit)
	return maxProfit
}

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
