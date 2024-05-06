package main

import (
	"fmt"
)

// minEatingSpeed finds the minimum rate at which Koko can consume all the
// bananas within a given time constraint.
func minEatingSpeed(piles []int, h int) int {
	return binarySearch(piles, h, 1, max(piles))
}

func binarySearch(piles []int, h, lo, hi int) int {
	if lo >= hi {
		return lo
	}
	mid := (hi + lo) / 2
	if canEatAll(piles, h, mid) {
		return binarySearch(piles, h, lo, mid)
	} else {
		return binarySearch(piles, h, mid+1, hi)
	}
}

func canEatAll(piles []int, h, k int) bool {
	time := 0
	for _, pile := range piles {
		time += (pile + k - 1) / k
	}
	return time <= h
}

func max(nums []int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
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
