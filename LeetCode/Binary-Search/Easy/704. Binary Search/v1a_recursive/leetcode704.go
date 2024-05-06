package main

import (
	"fmt"
)

// binarySearch uses a Helper Function to recursively search for the index of a target value
// within an ascending sorted slice.
func search(nums []int, target int) int {

	return binarySearch(nums, target, 0, len(nums)-1)
}

// binarySearch implements a recursive Binary Search Algorithm.
func binarySearch(nums []int, target, lo, hi int) int {
	if lo > hi {
		return -1
	}

	mid := lo + (hi-lo)/2

	switch {
	case target < nums[mid]:
		return binarySearch(nums, target, lo, mid-1)
	case target > nums[mid]:
		return binarySearch(nums, target, mid+1, hi)
	}
	return mid
}

///

func main() {

	//   0   1   2  3  4  5  6
	nums := []int{-3, -2, -1, 0, 1, 2, 3}
	fmt.Printf("nums: %v\n", nums)
	for target := -4; target <= 5; target++ {
		index := search(nums, target)
		fmt.Printf("target = %v \t index = %v\n", target, index)
	}

}
