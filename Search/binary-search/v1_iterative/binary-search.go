package main

import "fmt"

// binarySearch iteratively searches for a target value index within an
// ascending sorted slice; returns (-1) if it's not present.
func binarySearch(nums []int, target int) int {

	var lo, hi = 0, len(nums) - 1
	var mid int
	for lo <= hi {
		mid = (lo + hi) / 2
		switch {
		case target > nums[mid]:
			lo = mid + 1
		case target < nums[mid]:
			hi = mid - 1
		case target == nums[mid]:
			return mid
		}
	}
	return -1 // no solution
}

///

func main() { //   0   1   2  3  4  5  6
	nums := []int{-3, -2, -1, 0, 1, 2, 3}

	for _, target := range nums {
		index := binarySearch(nums, target)
		fmt.Printf("target = %2v : index = %v\n", target, index)
	}
}
