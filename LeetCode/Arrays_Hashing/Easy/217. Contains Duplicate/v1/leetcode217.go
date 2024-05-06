package main

import "fmt"

func containsDuplicate(nums []int) bool {
	cache := make(map[int]bool)

	for _, num := range nums {
		// read map(value) to check duplicate
		if cache[num] {
			return true
		} else {
			cache[num] = true
		}
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
