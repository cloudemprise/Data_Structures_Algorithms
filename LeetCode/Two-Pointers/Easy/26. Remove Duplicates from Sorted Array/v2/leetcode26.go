package main

import "fmt"

// Time Complexity = O(n)
// Space Complexity = O(1)

// removeDuplicates removes duplicate elements in-place from an ordered array
// and returns the number of unique elements.
func removeDuplicates(nums []int) int {
	var ln = len(nums) - 1
	var i = 0

	for j := 1; j <= ln && ln > 0; j++ {
		if nums[i] != nums[j] {
			nums[i+1] = nums[j]
			i++
		}
	}
	return i + 1
}

///

func main() {

	var nums []int

	nums = []int{1, 1, 2}
	fmt.Printf("Output: %v, nums = %v\n", removeDuplicates(nums), nums)

	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Printf("Output: %v, nums = %v\n", removeDuplicates(nums), nums)

	nums = []int{1, 2}
	fmt.Printf("Output: %v, nums = %v\n", removeDuplicates(nums), nums)

	nums = []int{1, 1}
	fmt.Printf("Output: %v, nums = %v\n", removeDuplicates(nums), nums)

	nums = []int{1}
	fmt.Printf("Output: %v, nums = %v\n", removeDuplicates(nums), nums)
}
