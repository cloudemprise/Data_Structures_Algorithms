package main

import "fmt"

func maximumWealth(accounts [][]int) int {

	var maxWealth int

	for _, customer := range accounts {
		var sum int

		for _, amount := range customer {
			sum += amount
		}
		if sum > maxWealth {
			maxWealth = sum
		}
	}
	return maxWealth
}

///

func main() {

	accounts := [][]int{
		{1, 2, 3},
		{3, 2, 1},
	}
	fmt.Printf("Accounts: %v\n", accounts)
	fmt.Printf("Richest Customer Wealth = %d", maximumWealth(accounts))

}
