package main

func productExceptSelf(nums []int) []int {

	products := make([]int, len(nums))
	// product := 1

	// brute force algo
	// for i := 0; i < len(nums); i++ {
	//     for j := 0; j < len(nums); j++ {
	//         if i != j {
	//             product *= nums[j]
	//         }
	//     }

	//     products[i] = product
	// }

	// O(n) solution
	prefix := 1
	for i := 0; i < len(nums); i++ {
		products[i] = prefix
		prefix *= nums[i]
	}

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		products[i] *= postfix
		postfix *= nums[i]
	}

	return products
}
