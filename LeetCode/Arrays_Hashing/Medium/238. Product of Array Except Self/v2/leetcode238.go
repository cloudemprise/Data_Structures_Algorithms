package main

import "fmt"

/*
This solution utilizes prefix and postfix multiplication to calculate the
products of all elements except the current element.

Prefix Products (Moving Forward):
	Initialize a variable leftProduct to 1.
	Iterate through the input array from left to right.
	For each element at index i, store the product of all elements to the left
		of it in the result array and update leftProduct by multiplying it with
		the current element.
	This step calculates the product of all elements to the left of each element.

Postfix Products (Moving Backward):
	Initialize a variable rightProduct to 1.
	Iterate through the input array from right to left.
	For each element at index i, multiply the previously calculated prefix
		product with the product of all elements to the right of it and store it
		in the result array.
	Update rightProduct by multiplying it with the current element.
	This step calculates the product of all elements to the right of each
		element and multiplies it with the corresponding prefix product.
*/

/*
Time and Space Complexity Analysis:

The time complexity of the above solution is O(n), where n is the number of
elements in the input array.

The space complexity is O(1) without considering the output array, and O(n)
when considering the output array.
*/

// productExceptSelf returns the product of all the input elements except the element at that index.
func productExceptSelf(nums []int) []int {
	ln := len(nums)
	// Space Complexity: O(n)
	result := make([]int, ln) // Need indices, preallocate.

	// Calculate the prefix products (moving forward):
	leftProduct := 1
	for i := 0; i < ln; i++ {
		result[i] = leftProduct
		leftProduct = leftProduct * nums[i]
	} // Time Complexity: O(n) / Space Complexity: O(1)

	// Calculate the postfix products (moving backward):
	rightProduct := 1
	for i := ln - 1; i >= 0; i-- {
		result[i] = result[i] * rightProduct
		rightProduct = rightProduct * nums[i]
	} // Time Complexity: O(n) / Space Complexity: O(1)

	return result
}

///

func main() {

	var nums []int

	nums = []int{1, 2, 3, 4}
	fmt.Printf("Input: %v\n", nums)
	fmt.Printf("Output: %v\n", productExceptSelf(nums))

	nums = []int{}
	fmt.Printf("Input: %v\n", nums)
	fmt.Printf("Output: %v\n", productExceptSelf(nums))

	nums = []int{1}
	fmt.Printf("Input: %v\n", nums)
	fmt.Printf("Output: %v\n", productExceptSelf(nums))

	nums = []int{0, 0, 0, 0}
	fmt.Printf("Input: %v\n", nums)
	fmt.Printf("Output: %v\n", productExceptSelf(nums))

}
