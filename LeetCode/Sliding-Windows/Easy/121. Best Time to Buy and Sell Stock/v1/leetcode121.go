package main

import (
	"fmt"
)

/*
NOT OPTIMAL!!!

Time Complexity = O(n^2) / Space Complexity = O(1)

This solution is a BRUTE FORCE approach to the "MaxProfit" problem in the
context of stock trading. It iterates through the list of stock prices and
calculates the maximum profit that can be made with one buy and one sell
operation. However, this solution has a time complexity of O(n^2), which is not
efficient for large input sizes.
*/

// maxProfit returns the maximum profit that can be made by buying and selling stocks.
func maxProfit(prices []int) int {

	maxProfit := 0
	for buy := 0; buy <= len(prices)-2; buy++ { // Can not buy on last day.
		for sell := buy + 1; sell <= len(prices)-1; sell++ {
			profit := prices[sell] - prices[buy]
			if profit > maxProfit {
				maxProfit = profit
			}
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
