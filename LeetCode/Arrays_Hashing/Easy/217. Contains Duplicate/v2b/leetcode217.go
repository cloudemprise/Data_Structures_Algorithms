package main

import "fmt"

/*
Complexity Analysis:
--------------------

The time complexity of this solution is O(n) because we iterate through the
input array once.

The space complexity is also O(n) because in the worst case, the map will
contain n unique elements.
*/

func containsDuplicate(nums []int) bool {
	// empty struct occupies no space
	cache := make(map[int]struct{})

	for _, num := range nums {
		// ok idiom to check for duplicates
		if _, ok := cache[num]; ok {
			return true
		}
		cache[num] = struct{}{} // empty struct flag
	}
	return false
}

///

func main() {
	var input []int

	input = []int{1, 2, 3, 1}
	fmt.Println("Want: true \t Got: ", containsDuplicate(input))

	input = []int{1, 2, 3, 4}
	fmt.Println("Want: false \t Got: ", containsDuplicate(input))

	input = []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println("Want: true \t Got: ", containsDuplicate(input))

	input = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1}
	fmt.Println("Want: true \t Got: ", containsDuplicate(input))

	input = []int{-9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1}
	fmt.Println("Want: true \t Got: ", containsDuplicate(input))
}
