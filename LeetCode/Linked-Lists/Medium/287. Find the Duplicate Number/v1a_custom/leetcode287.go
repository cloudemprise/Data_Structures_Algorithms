package main

import "fmt"

// Custom logic based on a modified insertion sort algorithm.
// https://en.wikipedia.org/wiki/Cycle_sort

/*
Time complexity = O(n^2)
Space complexity = O(1)
*/
/* func findDuplicate(nums []int) int {
	i := 0
	for i < len(nums) {

		// num is already sorted. Move forward.
		if nums[i] == i+1 {
			i++
			continue
		}

		// num at indexed-position is not sorted.
		if nums[i] != nums[nums[i]-1] {
			// Put num in its correct indexed-position.
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]

			// Check if swapped value is in its
			// correct indexed-position.
			if nums[i] == i+1 {
				i++ // Yes, move forward.
				continue
			}

		} else if nums[i] != i+1 {
			// num at indexed-position is already sorted
			// but num is incorrectly indexed.
			return nums[i] // Duplicate found.
		}

	}
	return -1
} */

func findDuplicate(nums []int) int {
	for nums[0] != nums[nums[0]] {
		nums[nums[0]], nums[0] = nums[0], nums[nums[0]]
		fmt.Println(nums)
	}

	return nums[0]
}

///

func main() {

	var input []int

	fmt.Println("Test1:")
	input = []int{1, 3, 4, 2, 2}
	fmt.Printf("input = %v\n", input)
	fmt.Printf("%v\n", findDuplicate(input))
	fmt.Println()

	//*-*

	fmt.Println("Test2:")
	input = []int{3, 1, 3, 4, 2}
	fmt.Printf("input = %v\n", input)
	fmt.Printf("%v\n", findDuplicate(input))
	fmt.Println()

	//*-*

	fmt.Println("Test3:")
	input = []int{8, 7, 1, 10, 17, 15, 18, 11, 16, 9, 19, 12, 5, 14, 3, 4, 2, 13, 18, 18}
	fmt.Printf("input = %v\n", input)
	fmt.Printf("%v\n", findDuplicate(input))
	fmt.Println()

	//*-*

}

//*-*
