package main

import (
	"fmt"
)

/*
The key idea is to determine which part of the mid-division is continuous and
then adjust the search range accordingly.

If the left part of the array is continuous, we check if the target is within
the sorted range. If it is, we update the `hi` pointer to `mid - 1` to search
the left part.

Otherwise, we update the `lo` pointer to `mid + 1` to search the right part. If
the right part of the array is continuous, we perform a similar check and update
the pointers accordingly.
*/

// Time Complexity  : O(log n)
// Space Complexity : O(1)

func search(nums []int, target int) int {

	lo, hi := 0, len(nums)-1
	for lo <= hi {

		// Divide the range.
		mid := (hi + lo) / 2

		if nums[mid] == target {
			return mid // break case.
		}

		// Determine which part of the mid-division is continuous.
		if nums[lo] <= nums[mid] {

			if target >= nums[lo] && target < nums[mid] {
				hi = mid - 1 // Shrink to fit.
			} else {
				lo = mid + 1 // target is not on this side.
			}

		} else { // nums[hi] > nums[mid]

			if target > nums[mid] && target <= nums[hi] {
				lo = mid + 1 // Shrink to fit.
			} else {
				hi = mid - 1 // target is not on this side.
			}
		}
	}
	return -1 // target not found.
}

///

func main() {

	var nums []int
	var target int

	var want int
	var got int

	//*-*

	fmt.Println("Test1:")
	nums = []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Printf("nums = %v\n", nums)
	target = 0
	fmt.Printf("target = %v\n", target)

	want = 4
	got = search(nums, target)
	fmt.Printf("Want: %v \n Got: %v\n", want, got)
	fmt.Println()

	//*-*

	/* fmt.Println("Test2:")
	nums = []int{1}
	fmt.Printf("nums = %v\n", nums)
	target = 1
	fmt.Printf("target = %v\n", target)

	want = 0
	got = search(nums, target)
	fmt.Printf("Want: %v \n Got: %v\n", want, got)
	fmt.Println() */

	//*-*

}
