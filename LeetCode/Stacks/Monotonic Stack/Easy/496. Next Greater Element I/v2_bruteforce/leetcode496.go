package main

import (
	"fmt"
)

// This version uses a brute force algorithm.

// Time Complexity  = O(n * m)
// Space Complexity = O(n)

// nextGreaterElement returns a list of the next greatest elements of nums1[i]
// that occur within the superset of nums2.
func nextGreaterElement(nums1 []int, nums2 []int) []int {

	// Record the index of each (unique) value.
	numIndex := make(map[int]int)
	for i, num := range nums2 {
		numIndex[num] = i
	}

	result := make([]int, len(nums1)) // preallocate for index management
	for i, current := range nums1 {

		result[i] = -1 // initialize default value

		// Brute force search for nextGreater.
		for j := numIndex[current] + 1; j < len(nums2); j++ {

			if nums2[j] > current {
				result[i] = nums2[j]
				break
			}
		} // Time: O(n)
	} // Time: O(m)
	return result
}

///

func main() {

	var nums1 []int
	var nums2 []int

	var want []int
	var got []int

	//*-*

	// This is a good example to run in debug mode to see the logic in action.
	fmt.Println("Test1:")
	nums1 = []int{3, 2, 5, 1, 4}
	fmt.Printf("nums1 = %v \n", nums1)
	nums2 = []int{4, 5, 2, 1, 3}
	fmt.Printf("nums2 = %v \n", nums2)

	want = []int{-1, 3, -1, 3, 5}
	got = nextGreaterElement(nums1, nums2)
	fmt.Printf("Want: %v \n Got: %v\n", want, got)
	fmt.Println()

	//*-*

	fmt.Println("Test2:")
	nums1 = []int{4, 1, 2}
	fmt.Printf("nums1 = %v \n", nums1)
	nums2 = []int{1, 3, 4, 2}
	fmt.Printf("nums2 = %v \n", nums2)

	want = []int{-1, 3, -1}
	got = nextGreaterElement(nums1, nums2)
	fmt.Printf("Want: %v \n Got: %v\n", want, got)
	fmt.Println()

	//*-*

	fmt.Println("Test3:")
	nums1 = []int{4, 1, 2}
	fmt.Printf("nums1 = %v \n", nums1)
	nums2 = []int{1, 2, 3, 4}
	fmt.Printf("nums2 = %v \n", nums2)

	want = []int{-1, 2, 3}
	got = nextGreaterElement(nums1, nums2)
	fmt.Printf("Want: %v \n Got: %v\n", want, got)
	fmt.Println()

	//*-*

	fmt.Println("Test4:")
	nums1 = []int{2, 1, 3}
	fmt.Printf("nums1 = %v \n", nums1)
	nums2 = []int{2, 3, 1}
	fmt.Printf("nums2 = %v \n", nums2)

	want = []int{3, -1, -1}
	got = nextGreaterElement(nums1, nums2)
	fmt.Printf("Want: %v \n Got: %v\n", want, got)
	fmt.Println()

	//*-*

}
