package main

import "fmt"

// Solution 1:
// Time Complexity : 3xO(n) = O(n)
// Space Complexity: 3xO(n) = O(n)
func productExceptSelf(nums []int) []int {
	ln := len(nums)
	// Space Complexity: O(n)
	result := make([]int, ln) // Need indices, preallocate.

	// Calculate the prefix products (moving forward):
	prefix := make([]int, ln)
	leftProduct := 1
	for i := 0; i < ln; i++ {
		prefix[i] = leftProduct
		leftProduct = leftProduct * nums[i]
	} // Time Complexity: O(n) / Space Complexity: O(n)

	// Calculate the postfix products (moving backward):
	postfix := make([]int, ln)
	rightProduct := 1
	for i := ln - 1; i >= 0; i-- {
		postfix[i] = rightProduct
		rightProduct = rightProduct * nums[i]
	} // Time Complexity: O(n) / Space Complexity: O(n)

	// Multiply pre & post products.
	for i := 0; i < ln; i++ {
		result[i] = prefix[i] * postfix[i]
	} // Time Complexity: O(n) / Space Complexity: O(1)

	return result
}

///

func main() {

	var nums []int

	nums = []int{1, 2, 3, 4}
	fmt.Printf("Input: %v\n", nums)
	fmt.Printf("Output: %v\n", productExceptSelf(nums))

}
