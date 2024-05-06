package main

import "fmt"

/*
Time Complexity = O(n)
Space Complexity = O(n)
*/

// Insert integer and its index into a map after check for remainder.
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		remainder := target - num
		if idx, ok := numMap[remainder]; ok { // Solution found.
			return []int{idx, i}
		}
		// Remainder not found, therefore
		// insert num:index into numMap[map]
		numMap[num] = i
	}
	return []int{} // No solution.
}

///

func main() {

	var input []int
	var target int

	input = []int{2, 7, 11, 15}
	target = 9
	fmt.Printf("\nInput = %v\n", input)
	fmt.Printf("target = %v\n", target)
	fmt.Println(twoSum(input, target))

	input = []int{3, 2, 4}
	target = 6
	fmt.Printf("\nInput = %v\n", input)
	fmt.Printf("target = %v\n", target)
	fmt.Println(twoSum(input, target))

	input = []int{3, 3}
	target = 6
	fmt.Printf("\nInput = %v\n", input)
	fmt.Printf("target = %v\n", target)
	fmt.Println(twoSum(input, target))

}
