package main

import (
	"fmt"
	"sort"
	"time"
)

// findKthLargest returns the Kth largest element of an unordered array.
func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

// quickselect (recursive) finds the Kth smallest element in an unordered array.
func quickSelect(nums []int, lo, hi, target int) int {
	if lo == hi {
		return nums[lo]
	}

	pivotIndex := hoarePartition(nums, lo, hi)

	if target == pivotIndex {
		return nums[target]
	} else if target < pivotIndex {
		return quickSelect(nums, lo, pivotIndex-1, target)
	} else {
		return quickSelect(nums, pivotIndex+1, hi, target)
	}
}

///

// hoarePartition partitions an unordered list about a chosen pivot element
// value and then returns the index of that pivot element after moving it into
// it final position.
func hoarePartition(nums []int, lo, hi int) int {

	// Select 1st element as pivot value.
	pivot := nums[lo]

	// Two pointers.
	left := lo + 1
	right := hi

	// do-while loop.
	var finished bool
	for !finished {

		// Search LHS.
		for left <= right && nums[left] <= pivot {
			left++ // Ignore numbers <= pivot.
		}

		// Search RHS.
		for nums[right] > pivot {
			right-- // Ignore numbers > pivot.
		}

		// Still within bounds?
		if left < right { // Swap elements.
			nums[left], nums[right] = nums[right], nums[left]
		}

		// Exit after pointers have passed each other,
		// thus final pivot index = right.
		finished = left > right
	}

	// Move pivot-value into its final position.
	nums[lo], nums[right] = nums[right], nums[lo]
	return right // Return pivot index
}

///

func main() {

	var nums []int
	var k, result int
	var start time.Time
	var elapsed time.Duration

	nums = []int{3, 2, 1, 5, 6, 4}
	k = 2
	fmt.Println("--- Unsorted --- \n", nums)
	fmt.Printf("k = %v\n", k)

	start = time.Now()
	result = findKthLargest(nums, k)
	elapsed = time.Since(start)

	fmt.Printf("Want: 5 Got: %v\n", result)
	fmt.Printf("Execution Time : %v\n", elapsed)

	sort.Ints(nums)
	fmt.Println("--- Sorted --- \n", nums)
	fmt.Println()

	///

	nums = []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k = 4
	fmt.Println("--- Unsorted --- \n", nums)
	fmt.Printf("k = %v\n", k)

	start = time.Now()
	result = findKthLargest(nums, k)
	elapsed = time.Since(start)

	fmt.Printf("Want: 4 Got: %v\n", result)
	fmt.Printf("Execution Time : %v\n", elapsed)

	sort.Ints(nums)
	fmt.Println("--- Sorted --- \n", nums)
	fmt.Println()

	///

	nums = []int{-1, 2, 0}
	k = 2
	fmt.Println("--- Unsorted --- \n", nums)
	fmt.Printf("k = %v\n", k)

	start = time.Now()
	result = findKthLargest(nums, k)
	elapsed = time.Since(start)

	fmt.Printf("Want: 0 Got: %v\n", result)
	fmt.Printf("Execution Time : %v\n", elapsed)

	sort.Ints(nums)
	fmt.Println("--- Sorted --- \n", nums)
	fmt.Println()

	///

}
