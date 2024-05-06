package main

import "fmt"

// This is just to illustrate what NOT to do as this algorithm assumes that the
// target value actually exists within the problem space, else the returned
// index is NOT valid. Rather use v2a_iterative.

// binarySearch searches iteratively for a target value index within an ascending sorted slice.
func binarySearch(nums []int, target int) int {

	var lo, hi = 0, len(nums) - 1
	var mid int
	for lo < hi {
		mid = (lo + hi) / 2
		if target > nums[mid] {
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

	fmt.Println("TEST1:")
	nums = []int{1, 2, 3, 4, 5}
	fmt.Printf("nums: \t\t %v\n", nums)
	target = 3
	fmt.Printf("target: \t %v \n", target)

	want = 2
	got = binarySearch(nums, target)
	fmt.Printf("Want: \t\t %v \n Got: \t\t %v \n", want, got)
	fmt.Println()

	///

	fmt.Println("TEST2:")
	nums = []int{1, 2, 4, 5}
	fmt.Printf("nums: \t\t %v\n", nums)
	target = 3
	fmt.Printf("target: \t %v \n", target)

	want = 2
	got = binarySearch(nums, target)
	fmt.Printf("Want: \t\t %v \n Got: \t\t %v \n", want, got)
	fmt.Println()

	///

	fmt.Println("TEST3:")
	nums = []int{1, 2, 3, 4, 5}
	fmt.Printf("nums: \t\t %v\n", nums)
	target = 6
	fmt.Printf("target: \t %v \n", target)

	want = 5
	got = binarySearch(nums, target)
	fmt.Printf("Want: \t\t %v \n Got: \t\t %v \n", want, got)
	fmt.Println()

}
