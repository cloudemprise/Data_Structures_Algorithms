package main

import (
	"fmt"
	"math/rand"
	"time"
)

// > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > >

/*
Selection Sort Algorithm
time complexity = O(n^2)
space complexity = O(1)

Characterised by swapping less elements than bubble sort but requiring
more comparisons.

An algorithm that repeatedly steps through a list and searches for the
index of the smallest value and then swaps that element with the next
smallest index of an array.

Each pass through the list will place the next smallest element value in its
proper place, thus: for i := 0; i < ln; i++ {...}
*/

// < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < <

// Iterative Selection Sort Algorithm
// Time Complexity: O(n^2)
// Space Complexity: O(1)
// Swaps = O(n)
func selectionSortIterative(nums []int) {
	ln := len(nums)
	for i := 0; i < ln; i++ {
		minIdx := i
		for j := i; j < ln; j++ {
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}
}

// Linear Tail Recursion Selection Sort
// Time Complexity: O(n^2)
// Space Complexity: O(1)
// Swaps = O(n)
func selectionSortRecursive(nums []int) {
	ln := len(nums)
	if ln <= 1 {
		return
	} else {
		var minIdx int
		for i := 1; i < ln; i++ {
			if nums[i] < nums[minIdx] {
				minIdx = i
			}
		}
		nums[0], nums[minIdx] = nums[minIdx], nums[0]
		selectionSortRecursive(nums[1:])
	}
}

//-------------------------------------

func main() {

	var nums []int
	var start time.Time
	var elapsed time.Duration

	nums = generateSlice(9)
	fmt.Println("Iterative:")
	fmt.Println("--- Unsorted --- \n", nums)
	start = time.Now()
	selectionSortIterative(nums)
	elapsed = time.Since(start)
	fmt.Println("--- Sorted ---\n", nums)
	fmt.Printf("Execution Time : %v\n", elapsed)
	fmt.Println()

	//*-*

	nums = generateSlice(9)
	fmt.Println("Recursive:")
	fmt.Println("--- Unsorted --- \n", nums)
	start = time.Now()
	selectionSortRecursive(nums)
	elapsed = time.Since(start)
	fmt.Println("--- Sorted ---\n", nums)
	fmt.Printf("Execution Time : %v\n", elapsed)
	fmt.Println()

	//*-*

}

//*-*

// generate a slice of size, size filled with random numbers
func generateSlice(size int) []int {
	nums := make([]int, size)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {
		// for negative numbers
		nums[i] = r.Intn(999) - r.Intn(999)
	}
	return nums
}

//*-*
