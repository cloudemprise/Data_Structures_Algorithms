package main

import "fmt"

// Time: O(n) / Space: O(1)

// maxProduct ...
func maxProduct(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	// 1st & 2nd maximum values.
	var first, second int

	// Build window-span.
	if nums[0] > nums[1] {
		first, second = nums[0], nums[1]
	} else {
		first, second = nums[1], nums[0]
	}

	// Implement sliding window, keeping track of
	// the two largest numbers.
	for i := 2; i < len(nums); i++ {

		if nums[i] > second {
			second = nums[i]
			if second > first {
				first, second = second, first
			}
		}
	}
	return (first - 1) * (second - 1)
}

/* // A more concise version.
func maxProduct(nums []int) int {

	var first, second int

	for _, num := range nums {
		if num > first {
			second = first // keep copy of next max
			first = num
		} else if num > second {
			second = num
		}
	}
	return (first - 1) * (second - 1)
} */

///

func main() {

	input := []int{3, 4, 5, 2}
	want := 12
	got := maxProduct(input)
	fmt.Println("Want: ", want, "Got: ", got)

}
