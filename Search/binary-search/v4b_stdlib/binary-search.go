package main

import (
	"fmt"
	"sort"
)

// binarySearch uses the stdlib function 'sort.SearchInts()' to perform a binary search
// for the index of a target value within an ascending sorted slice.
func binarySearch(nums []int, target int) int {

	index := sort.SearchInts(nums, target)

	/* if index >= len(nums) || nums[index] != target {
		return -1 // target not found
	}
	return index */

	if index < len(nums) && nums[index] == target {
		return index
	}
	return -1 // target not found
}

///

func main() {

	//   0   1   2  3  4  5  6
	nums := []int{-3, -2, -1, 0, 1, 2, 3}
	fmt.Printf("nums: %v\n", nums)
	for target := -4; target <= 5; target++ {
		index := binarySearch(nums, target)
		fmt.Printf("target = %v \t index = %v\n", target, index)
	}

}
