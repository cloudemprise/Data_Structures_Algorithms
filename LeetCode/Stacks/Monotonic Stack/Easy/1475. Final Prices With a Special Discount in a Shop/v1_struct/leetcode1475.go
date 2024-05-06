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

// We will see later that the price value needs not to be remembered.
type priceIndex struct {
	price int
	index int
}

// finalPrices returns the discounted price a customer will pay per item.
func finalPrices(prices []int) []int {

	// In order for there to be a discount available, we need to find prices
	// on the stack that are GREATER_OR_EQUAL the currentPrice, therefore use an
	// Increasing Monotonic Stack:

	var stack []priceIndex // Remember price & index.

	for i, currentPrice := range prices {
		// For each stack element GREATER_OR_EQUAL to currentPrice.
		for (len(stack) > 0) && (currentPrice <= stack[len(stack)-1].price) {

			// Calculate discount & overwrite []prices (saves space).
			prices[stack[len(stack)-1].index] = stack[len(stack)-1].price - currentPrice

			// The element at the top of stack has now been solved for and no
			// longer contributes to further solutions, thus pop() it off and
			// discard it.
			stack = stack[:len(stack)-1]
		}
		// Push current price & index.
		stack = append(stack, priceIndex{
			price: currentPrice,
			index: i,
		})
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
