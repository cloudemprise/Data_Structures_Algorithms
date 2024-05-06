package main

import "fmt"

// This algorithm assumes that the target value actually exists within the
// problem space, else the returned index is not valid.

// binarySearch calls a Helper Function that recursively searches for a target
// value index within an ascending sorted slice.
func binarySearch(nums []int, target int) int {

	return search(nums, target, 0, len(nums)-1)
}

// search is a Helper Function that implements a recursive binary search.
func search(nums []int, target, lo, hi int) int {
	if lo > hi {
		return -1
	}

	mid := (hi + lo) / 2

	switch {
	case target > nums[mid]:
		return search(nums, target, mid+1, hi)
	case target < nums[mid]:
		return search(nums, target, lo, mid-1)
	}
	return mid // tail recursion
}

///

func main() { //   0   1   2  3  4  5  6
	nums := []int{-3, -2, -1, 0, 1, 2, 3}

	for _, target := range nums {
		index := binarySearch(nums, target)
		fmt.Printf("target = %2v : index = %v\n", target, index)
	}
}
