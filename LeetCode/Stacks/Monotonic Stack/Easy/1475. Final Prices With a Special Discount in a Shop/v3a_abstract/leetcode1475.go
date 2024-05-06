package main

import (
	"fmt"
)

/*
Overview:

Implement "The Next Smallest Element" algorithm by making use of an Increasing
Monotonic Stack.
*/

// Time Complexity : O(n)
// Space Complexity: O(n)

// finalPrices returns the discounted price a customer will pay per item.
func finalPrices(prices []int) []int {

	// In order for there to be a discount available, we need to find prices
	// on the stack that are GREATER_OR_EQUAL the currentPrice, therefore use an
	// Increasing Monotonic Stack:

	indexStack := newStack()

	for index, currentPrice := range prices {
		// For each stack element GREATER_OR_EQUAL to currentPrice.
		for (indexStack.length() > 0) && (currentPrice <= prices[indexStack.peek()]) {

			// Calculate discount & overwrite []prices (saves space).
			prices[indexStack.peek()] -= currentPrice

			// The element at the top of stack has now been solved for and no
			// longer contributes to further solutions, thus pop() it off and
			// discard it.
			indexStack.pop()
		}
		// Push current index.
		indexStack.push(index)
	}
	return prices
}

///

func main() {

	var prices []int

	var want []int
	var got []int

	///

	fmt.Println("Test1:")
	prices = []int{8, 4, 6, 2, 3}
	fmt.Printf("prices = %v \n", prices)

	want = []int{4, 2, 4, 2, 3}
	got = finalPrices(prices)
	fmt.Printf("Want: %v \n Got: %v\n", want, got)
	fmt.Println()

	///

	fmt.Println("Test2:")
	prices = []int{1, 2, 3, 4, 5}
	fmt.Printf("prices = %v \n", prices)

	want = []int{1, 2, 3, 4, 5}
	got = finalPrices(prices)
	fmt.Printf("Want: %v \n Got: %v\n", want, got)
	fmt.Println()

	///

	fmt.Println("Test3:")
	prices = []int{10, 1, 1, 6}
	fmt.Printf("prices = %v \n", prices)

	want = []int{9, 0, 1, 6}
	got = finalPrices(prices)
	fmt.Printf("Want: %v \n Got: %v\n", want, got)
	fmt.Println()

	///

	fmt.Println("Test4:")
	prices = []int{8, 7, 4, 2, 8, 1, 7, 7, 10, 1}
	fmt.Printf("prices = %v \n", prices)

	want = []int{1, 3, 2, 1, 7, 0, 0, 6, 9, 1}
	got = finalPrices(prices)
	fmt.Printf("Want: %v \n Got: %v\n", want, got)
	fmt.Println()
}
