package main

func twoSum(nums []int, target int) []int {
	// brute force solution O(n^2)
	// for i := 0; i < len(nums); i++ {
	//     for j := i + 1; j < len(nums); j++ {
	//         if nums[i] + nums[j] == target {
	//             return []int{i, j}
	//         }
	//     }
	// }
	// return nil

	// O(n)
	seen := make(map[int]int) // value -> index
	for i, num := range nums {
		difference := target - num

		if value, exists := seen[difference]; exists {
			return []int{value, i}
		}

		seen[num] = i
	}

	return nil
}
