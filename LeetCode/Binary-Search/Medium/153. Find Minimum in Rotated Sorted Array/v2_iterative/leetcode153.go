package main

import (
	"fmt"
)

/*
The difference between the standard algorithm and this is that here we are
searching for a MINIMUM value as opposed to a TARGET value. You can think of
the target value being in the middle of the problem set because of the rotation.
*/

// Time Complexity  : O(log n)
// Space Complexity : O(1)

// findMin finds the MINIMUM value in a rotated sorted array
// using indices comparisons.
func findMin(nums []int) int {

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

	return nums[lo]
}

///

func main() {

	var nums []int

	var want int
	var got int

	///

	fmt.Println("Test1:")
	nums = []int{3, 4, 5, 1, 2}
	fmt.Printf("nums = %v\n", nums)

	want = 1
	got = findMin(nums)
	fmt.Printf("Want: %v \n Got: %v\n", want, got)
	fmt.Println()

	///

	fmt.Println("Test2:")
	nums = []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Printf("nums = %v\n", nums)

	want = 0
	got = findMin(nums)
	fmt.Printf("Want: %v \n Got: %v\n", want, got)
	fmt.Println()

	///

	fmt.Println("Test3:")
	nums = []int{11, 13, 15, 17}
	fmt.Printf("nums = %v\n", nums)

	want = 11
	got = findMin(nums)
	fmt.Printf("Want: %v \n Got: %v\n", want, got)
	fmt.Println()

	///

	fmt.Println("Test4:")
	nums = []int{2, 1}
	fmt.Printf("nums = %v\n", nums)

	want = 1
	got = findMin(nums)
	fmt.Printf("Want: %v \n Got: %v\n", want, got)
	fmt.Println()

	///

	fmt.Println("Test5:")
	nums = []int{3, 1, 2}
	fmt.Printf("nums = %v\n", nums)

	want = 1
	got = findMin(nums)
	fmt.Printf("Want: %v \n Got: %v\n", want, got)
	fmt.Println()

}
