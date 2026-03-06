package main

import (
	"fmt"
)

func main() {
	// nums1 := []int{3,4,5,6}
	// nums2 := []int{4, 5, 6}

	// result := twoSum(nums2, 10)
	// fmt.Println("result:", result)

	// strs := []string{"act", "pots", "tops", "cat", "stop", "hat"}
	// result := groupAnagrams(strs)
	// fmt.Println("result:", result)

	// elements := []int{1, 2, 2, 3, 3, 3}
	// // elements := []int{1,1,1,2,2,3}
	// k := 2
	// result := topKFrequent(elements, k)
	// fmt.Printf("result: %v\n", result)

	elements := []int{1, 2, 4, 6}
	result := productExceptSelf(elements)
	fmt.Printf("result: %v\n", result)

	// var s []string
	// s = append(s, "a")
	// if s == nil {
	// 	fmt.Println("s == nil")
	// } else {
	// 	fmt.Println("s:", s)
	// }

}
