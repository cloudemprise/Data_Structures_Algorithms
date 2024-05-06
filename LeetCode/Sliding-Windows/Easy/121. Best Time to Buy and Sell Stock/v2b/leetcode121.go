package main

import (
	"fmt"
	"math"
)

/*
Time Complexity = O(n) / Space Complexity = O(1)

This solution iterates through the array of prices only once, keeping track of
the minimum price and the maximum profit found so far. Therefore, its time
complexity is linear with respect to the size of the input array.
*/

// maxProfit returns the maximum profit that can be made by buying and selling stocks.
func maxProfit(prices []int) int {
	minPrice := math.MaxInt32
	maxProfit := 0
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}
	return maxProfit
}

///

func main() {

	var input []int

	input = []int{7, 1, 5, 3, 6, 4}
	fmt.Printf("Prices = %v\n", input)
	fmt.Printf("Want: 5 \t Got: %v\n", maxProfit(input))
	fmt.Println()

	input = []int{7, 6, 4, 3, 1}
	fmt.Printf("Prices = %v\n", input)
	fmt.Printf("Want: 0 \t Got: %v\n", maxProfit(input))
	fmt.Println()

	input = []int{2, 1}
	fmt.Printf("Prices = %v\n", input)
	fmt.Printf("Want: 0 \t Got: %v\n", maxProfit(input))
	fmt.Println()

	input = []int{1, 2}
	fmt.Printf("Prices = %v\n", input)
	fmt.Printf("Want: 1 \t Got: %v\n", maxProfit(input))
	fmt.Println()

	input = []int{1}
	fmt.Printf("Prices = %v\n", input)
	fmt.Printf("Want: 0 \t Got: %v\n", maxProfit(input))
	fmt.Println()

	input = []int{1, 2, 4}
	fmt.Printf("Prices = %v\n", input)
	fmt.Printf("Want: 3 \t Got: %v\n", maxProfit(input))
	fmt.Println()

	input = []int{5, 2, 7, 1, 4, 3}
	fmt.Printf("Prices = %v\n", input)
	fmt.Printf("Want: 5 \t Got: %v\n", maxProfit(input))
	fmt.Println()

}
