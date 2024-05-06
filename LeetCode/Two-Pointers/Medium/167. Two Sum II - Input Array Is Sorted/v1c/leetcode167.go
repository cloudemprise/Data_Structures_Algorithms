package main

import "fmt"

// Time Complexity  = O(n)
// Space Complexity = O(1)

// twoSum finds two elements within a sorted array such that they add up to a
// specific target number using the two-pointer approach.
func twoSum(numbers []int, target int) []int {

	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		switch {
		case sum > target:
			right--
		case sum < target:
			left++
		default:
			return []int{left + 1, right + 1}
		}
	}
	return []int{}
}

///

func main() {

	var input []int
	var target int

	input = []int{2, 3, 5, 7}
	target = 9
	fmt.Printf("Input: %v\n", input)
	fmt.Printf("Target: %v\n", target)
	fmt.Printf("Want: [1 4] \t Got: %v\n", twoSum(input, target))
	fmt.Println()

	///

	input = []int{2, 7, 11, 15}
	target = 9
	fmt.Printf("Input: %v\n", input)
	fmt.Printf("Target: %v\n", target)
	fmt.Printf("Want: [1 2] \t Got: %v\n", twoSum(input, target))
	fmt.Println()

	///

	input = []int{2, 3, 4}
	target = 6
	fmt.Printf("Input: %v\n", input)
	fmt.Printf("Target: %v\n", target)
	fmt.Printf("Want: [1 3] \t Got: %v\n", twoSum(input, target))
	fmt.Println()

	///

	input = []int{5, 25, 75}
	target = 100
	fmt.Printf("Input: %v\n", input)
	fmt.Printf("Target: %v\n", target)
	fmt.Printf("Want: [2 3] \t Got: %v\n", twoSum(input, target))
	fmt.Println()

	///

}
