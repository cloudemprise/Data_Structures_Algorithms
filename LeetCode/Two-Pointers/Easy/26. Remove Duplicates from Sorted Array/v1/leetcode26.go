package main

import "fmt"

// Time Complexity = O(n)
// Space Complexity = O(1)

// removeDuplicates removes duplicate elements in-place from an ordered array
// and returns the number of unique elements.
func removeDuplicates(nums []int) int {
	var ln = len(nums) - 1
	var i, j = 0, 1

	for ln > 0 && j <= ln {
		switch {
		case nums[i] != nums[j]:
			nums[i+1] = nums[j]
			i++
			j++
		case nums[i] == nums[j]:
			j++
		default:
			panic("invalid state")
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
