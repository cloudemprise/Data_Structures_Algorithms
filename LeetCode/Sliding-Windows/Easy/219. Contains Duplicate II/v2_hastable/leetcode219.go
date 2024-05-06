package main

import "fmt"

// Time Complexity = O(n) / Space Complexity = O(n)

// Store each num in map and their corresponding indices, i.e. map(num)index

// containsNearbyDuplicate ...
func containsNearbyDuplicate(nums []int, k int) bool {

	// Create map to store values and their indices.
	cache := make(map[int]int)

	for i, num := range nums {

		// Check if value has been seen before.
		if j, ok := cache[num]; ok {
			// If so, check if index(j) is within range of k
			if (i - j) <= k {
				return true
			}
		}
		// Remember index of last seen value.
		cache[num] = i
	}
	return false // No luck.
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
