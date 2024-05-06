package main

import (
	"fmt"
)

// binarySearch iteratively searches for the index of a target value within an
// ascending sorted slice.
func search(nums []int, target int) int {

	lo, hi := 0, len(nums)-1
	var mid int

	for lo <= hi {

		mid = (hi + lo) / 2

		switch {
		case target < nums[mid]:
			hi = mid - 1
		case target > nums[mid]:
			lo = mid + 1
		case target == nums[mid]:
			return mid
		}
	}
	return -1 // target not found
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
