package main

import "fmt"

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
