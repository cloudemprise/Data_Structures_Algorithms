package main

import "fmt"

// Time Complexity O(n^2) : Space Complexity O(1)

// Suboptimum. Double loop with out-of-bounds break.

// containsNearbyDuplicate ...
func containsNearbyDuplicate(nums []int, k int) bool {

	for i := 0; i < len(nums)-1; i++ {

		// Only loop as far as the window-span.
		for j := 1; j <= k; j++ {
			// Check if loop indices extend beyond end of input list.
			if i+j > len(nums)-1 {
				break
			} else if nums[i] == nums[i+j] {
				return true
			}
		}
	}
	return false
}

///

func main() {
	var input []int
	var k int

	input = []int{1, 2, 3, 1}
	k = 3
	fmt.Printf("Input = %v\n", input)
	fmt.Printf("k = %v\n", k)
	fmt.Printf("Want: true \t Got: %v\n", containsNearbyDuplicate(input, k))
	fmt.Println()

	input = []int{1, 0, 1, 1}
	k = 1
	fmt.Printf("Input = %v\n", input)
	fmt.Printf("k = %v\n", k)
	fmt.Printf("Want: true \t Got: %v\n", containsNearbyDuplicate(input, k))
	fmt.Println()

	input = []int{1, 2, 3, 1, 2, 3}
	k = 2
	fmt.Printf("Input = %v\n", input)
	fmt.Printf("k = %v\n", k)
	fmt.Printf("Want: false \t Got: %v\n", containsNearbyDuplicate(input, k))
	fmt.Println()

	input = []int{1, 2, 3, 1, 2, 3, 2}
	k = 2
	fmt.Printf("Input = %v\n", input)
	fmt.Printf("k = %v\n", k)
	fmt.Printf("Want: true \t Got: %v\n", containsNearbyDuplicate(input, k))
	fmt.Println()

	input = []int{99, 99}
	k = 2
	fmt.Printf("Input = %v\n", input)
	fmt.Printf("k = %v\n", k)
	fmt.Printf("Want: true \t Got: %v\n", containsNearbyDuplicate(input, k))
	fmt.Println()

	input = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9}
	k = 3
	fmt.Printf("Input = %v\n", input)
	fmt.Printf("k = %v\n", k)
	fmt.Printf("Want: true \t Got: %v\n", containsNearbyDuplicate(input, k))
	fmt.Println()

	input = []int{1, 2, 3, 4, 5}
	k = 3
	fmt.Printf("Input = %v\n", input)
	fmt.Printf("k = %v\n", k)
	fmt.Printf("Want: false \t Got: %v\n", containsNearbyDuplicate(input, k))
	fmt.Println()

}
