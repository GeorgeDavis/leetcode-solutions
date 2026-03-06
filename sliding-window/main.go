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
