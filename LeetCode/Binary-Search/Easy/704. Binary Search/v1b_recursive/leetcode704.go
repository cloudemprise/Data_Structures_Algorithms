package main

import (
	"fmt"
)

// search is a recursive closure solution to leetcode 704.

// search implements a recursive closure to perform the Binary Search Algorithm
// that searches for the index of a target value within an ascending sorted slice.
func search(nums []int, target int) int {

	// Using a Closure:
	var binarySearch func(int, int) int
	binarySearch = func(lo, hi int) int {
		if lo > hi {
			return -1
		}

		mid := (hi + lo) / 2

		switch {
		case target < nums[mid]:
			return binarySearch(lo, mid-1)
		case target > nums[mid]:
			return binarySearch(mid+1, hi)
		}
		return mid // tail recursion
	}

	return binarySearch(0, len(nums)-1) // closure call
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
