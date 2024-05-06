package main

import "fmt"

// searchInsert searches iteratively for a target value index within an ascending
// sorted slice. If the target value  doesn't exist then the return value is the
// index at which the target value would be if it were inserted in order.
func searchInsert(nums []int, target int) int {

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
	got = searchInsert(nums, target)
	fmt.Printf("Want: \t\t %v \n Got: \t\t %v \n", want, got)
	fmt.Println()

	///

	fmt.Println("TEST2:")
	nums = []int{1, 2, 4, 5}
	fmt.Printf("nums: \t\t %v\n", nums)
	target = 3
	fmt.Printf("target: \t %v \n", target)

	want = 2
	got = searchInsert(nums, target)
	fmt.Printf("Want: \t\t %v \n Got: \t\t %v \n", want, got)
	fmt.Println()

	///

	fmt.Println("TEST3:")
	nums = []int{1, 2, 3, 4, 5}
	fmt.Printf("nums: \t\t %v\n", nums)
	target = 6
	fmt.Printf("target: \t %v \n", target)

	want = 5
	got = searchInsert(nums, target)
	fmt.Printf("Want: \t\t %v \n Got: \t\t %v \n", want, got)
	fmt.Println()

}
