package main

import (
	"fmt"
	"math"
)

// Time Complexity : O(n + log(biggestPile) * n) = O(n log(biggestPile))

// Space Complexity: O(1)

// minEatingSpeed finds the minimum rate at which Koko can consume all the
// bananas within a given time constraint.
func minEatingSpeed(piles []int, h int) int {

	// Koko can not eat more per hour than the biggest pile, so use this as
	// an arbitrary upper limit with which to search for the minimum rate.
	biggestPile := 0
	for _, pile := range piles {
		if pile > biggestPile {
			biggestPile = pile
		}
	} // Time: O(n)

	// Suggested content as a replacement for the above.
	// biggestPile := slices.Max(piles)

	// Binary Search Algorithm:
	lo, hi := 1, biggestPile
	var rate int
	for lo < hi {

		rate = (lo + hi) / 2 // binary bisection.

		// Each pile can only be eaten in hourly increments. Based on a
		// chosen rate calculate the time taken to eat all bananas.
		var timeTaken int
		for _, pile := range piles {
			timeTaken += int(math.Ceil((float64(pile) / float64(rate))))
		} // Time: O(n)

		// Search for minimum rate, i.e. min(lo)
		if timeTaken > h { // too slow
			lo = rate + 1 // shift window right.
		} else { // not a minimum
			hi = rate // shift window left.
		}

	} // Time: O(log(biggestPile))

	return lo // minimum rate that satisfies constraints.
}

///

func main() {

	var piles []int
	var hours int

	var want int
	var got int

	//*-*

	fmt.Println("Test1:")
	piles = []int{3, 6, 7, 11}

	fmt.Printf("piles = %v\n", piles)
	hours = 8
	fmt.Printf("hours = %v\n", hours)

	want = 4
	got = minEatingSpeed(piles, hours)
	fmt.Printf("Want: %v \n Got: %v\n", want, got)
	fmt.Println()

	//*-*

	fmt.Println("Test2:")
	piles = []int{30, 11, 23, 4, 20}

	fmt.Printf("piles = %v\n", piles)
	hours = 5
	fmt.Printf("hours = %v\n", hours)

	want = 30
	got = minEatingSpeed(piles, hours)
	fmt.Printf("Want: %v \n Got: %v\n", want, got)
	fmt.Println()

	//*-*

	fmt.Println("Test3:")
	piles = []int{312884470}

	fmt.Printf("piles = %v\n", piles)
	hours = 312884469
	fmt.Printf("hours = %v\n", hours)

	want = 2
	got = minEatingSpeed(piles, hours)
	fmt.Printf("Want: %v \n Got: %v\n", want, got)
	fmt.Println()

	//*-*

}
