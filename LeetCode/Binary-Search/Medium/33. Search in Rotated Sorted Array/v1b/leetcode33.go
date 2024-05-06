package main

import (
	"fmt"
)

// Time Complexity  : O(log n)
// Space Complexity : O(1)

// search finds the index of a target value in a rotated sorted array.
func search(nums []int, target int) int {

	// 1. Find the index (split) of the minimum element.
	split := findMinIndex(nums)

	// 2. Designate boundary conditions according to split.
	lo, hi := 0, len(nums)-1
	if target <= nums[len(nums)-1] { // target in upper section
		lo = split
	} else {
		hi = split - 1 // target in lower section
	}

	// 3. Perform a target binary search.
	var mid int
	for lo <= hi {

		mid = (hi + lo) / 2

		switch {
		case target > nums[mid]:
			lo = mid + 1
		case target < nums[mid]:
			hi = mid - 1
		default:
			return mid // target found.
		}
	}
	return -1 // target not found.
}

///

// binarySearch iteratively searches for the index of a target value within an
// ascending sorted slice.
func binarySearch(nums []int, target int) int {

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

// findMinIndex finds the index of the MINIMUM value in a rotated sorted array
// using indices comparisons.
func findMinIndex(nums []int) int {

	lo, hi := 0, len(nums)-1
	var mid int

	for lo < hi {

		mid = (hi + lo) / 2

		if nums[mid] > nums[hi] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

///

func main() {

	var nums []int
	var target int

	var want int
	var got int

	///

	nums = []int{4, 5, 1, 2, 3}
	target = 1
	want = 2

	fmt.Println("TEST1:")
	fmt.Printf("nums: \t %v \n", nums)
	fmt.Printf("target:  %v \n", target)
	got = search(nums, target)
	fmt.Printf("Want: \t %v \n Got: \t %v \n", want, got)
	fmt.Println()

	///

	nums = []int{4, 5, 1, 2, 3}
	target = 4
	want = 0

	fmt.Println("TEST2:")
	fmt.Printf("nums: \t %v \n", nums)
	fmt.Printf("target:  %v \n", target)
	got = search(nums, target)
	fmt.Printf("Want: \t %v \n Got: \t %v \n", want, got)
	fmt.Println()

	///

	nums = []int{4, 5, 1, 2, 3}
	target = 3
	want = 4

	fmt.Println("TEST3:")
	fmt.Printf("nums: \t %v \n", nums)
	fmt.Printf("target:  %v \n", target)
	got = search(nums, target)
	fmt.Printf("Want: \t %v \n Got: \t %v \n", want, got)
	fmt.Println()

	///

	nums = []int{8, 10, 2, 4, 6}
	target = 1
	want = -1

	fmt.Println("TEST4:")
	fmt.Printf("nums: \t %v \n", nums)
	fmt.Printf("target:  %v \n", target)
	got = search(nums, target)
	fmt.Printf("Want: \t %v \n Got: \t %v \n", want, got)
	fmt.Println()

	///

	nums = []int{12, 15, 3, 6, 9}
	target = 11
	want = -1

	fmt.Println("TEST5:")
	fmt.Printf("nums: \t %v \n", nums)
	fmt.Printf("target:  %v \n", target)
	got = search(nums, target)
	fmt.Printf("Want: \t %v \n Got: \t %v \n", want, got)
	fmt.Println()

	///

	nums = []int{12, 15, 3, 6, 9}
	target = 10
	want = -1

	fmt.Println("TEST6:")
	fmt.Printf("nums: \t %v \n", nums)
	fmt.Printf("target:  %v \n", target)
	got = search(nums, target)
	fmt.Printf("Want: \t %v \n Got: \t %v \n", want, got)
	fmt.Println()

	///

}
