package main

import (
	"fmt"
	"math/rand"
	"time"
)

// > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > >

/*
Quicksort Algorithm

*/

// < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < <

// Quicksort Algorithm using Lomuto partition scheme.
// Ref:
// https://www.golangprograms.com/data-structure-and-algorithms/golang-program-for-implementation-of-quick-sort.html
// Time Complexity: Average = O(nlogn) Worst = O(n^2)
// Space Complexity: O(1)
// Swaps ~ 3x more than Hoare.
func quickSortLomuto(nums []int) {
	ln := len(nums)

	if ln < 2 { // base case.
		return
	}

	// Two pointers.
	lo, hi := 0, ln-1

	// Select a random index => pivot-value.
	pivot := rand.Intn(ln)
	// Move pivot-value to end of slice.
	nums[pivot], nums[hi] = nums[hi], nums[pivot]

	// Search for & move values less than pivot-value down.
	for i := 0; i < hi; i++ {
		if nums[i] < nums[hi] {
			nums[lo], nums[i] = nums[i], nums[lo]
			lo++ // move ptr forward after swap.
		}
	}
	// Move pivot-value into its proper position.
	nums[lo], nums[hi] = nums[hi], nums[lo]

	// Divide & Conquer.
	// Recursive calls exclude pivot-value.
	quickSortLomuto(nums[:lo])
	quickSortLomuto(nums[lo+1:])
}

//-------------------------------------

// Quicksort Algorithm using Hoare partition scheme.
// Ref:
// https://en.wikipedia.org/wiki/Quicksort#Hoare_partition_scheme
// https://www.geeksforgeeks.org/hoares-vs-lomuto-partition-scheme-quicksort/
// Time Complexity: Average = O(nlogn) Worst = O(n^2)
// Space Complexity: O(1)
// Swaps ~ 3x less than Lomuto.
func quickSortHoare1(nums []int, lo, hi int) {

	if lo < hi {
		pivot := partitionHoare1(nums, lo, hi)
		// Divide & Conquer.
		quickSortHoare1(nums, lo, pivot) // Note pivot included.
		quickSortHoare1(nums, pivot+1, hi)
	}
}

//*-*

func partitionHoare1(nums []int, lo, hi int) int {

	// element from middle of slice
	//pivot := nums[(hi-lo)/2+lo]
	pivot := nums[lo]

	l := lo - 1 // left index
	r := hi + 1 // right index

	for { // loop forever

		l++ // move left index right at least once.
		for nums[l] < pivot {
			l++ // search for next element > pivot.
		}

		r-- // move right index left at least once.
		for nums[r] > pivot {
			r-- // search for next element < pivot.
		}

		// if indices cross return.
		if l >= r {
			return r
		}

		// swap elements at left & right indices.
		nums[l], nums[r] = nums[r], nums[l]
	}
}

//-------------------------------------

// Quicksort Algorithm using Hoare2 partition scheme.
// Ref: Karumanchi - DSA Thinking in Go. Pg. 324
// Time Complexity: Average = O(nlogn) Worst = O(n^2)
// Space Complexity: O(1)
// Swaps ~ 3x less than Lomuto.
func quickSortHoare2(nums []int, lo, hi int) {

	if lo < hi { // exit if less than one element.
		pivot := partitionHoare2(nums, lo, hi)
		// Divide & Conquer. Note not pivot included.
		quickSortHoare2(nums, lo, pivot-1)
		quickSortHoare2(nums, pivot+1, hi)
	}
}

//*-*

func partitionHoare2(nums []int, lo, hi int) int {

	for lo < hi {

		for lo < hi && nums[lo] <= nums[hi] {
			hi--
		}
		if lo < hi {
			nums[lo], nums[hi] = nums[hi], nums[lo]
			lo++
		}

		for lo < hi && nums[lo] <= nums[hi] {
			lo++
		}
		if lo < hi {
			nums[lo], nums[hi] = nums[hi], nums[lo]
			hi--
		}
	}
	return lo
}

//-------------------------------------

// My Recursive Variant:
// Quicksort Algorithm using Hoare partition scheme.
func myQuickSort(nums []int, lo, hi int) {

	if lo < hi { // exit if less than one element.
		pivot := myPartition(nums, lo, hi)
		// Divide & Conquer. Note not pivot included.
		myQuickSort(nums, lo, pivot-1)
		myQuickSort(nums, pivot+1, hi)
	}
}

// My Iterative Variant2:
// Based on Hoare's Partitioning Method.
func myPartition(nums []int, lo, hi int) int {

	// Select 1st element as pivot value.
	pivot := nums[lo]

	// Create two pointers.
	l := lo + 1
	r := hi

	for l <= r {

		// Search left-hand side.
		// Skip numbers equal to pivot.
		// Don't allow lhs to move more than one passed rhs.
		for l <= r && nums[l] <= pivot {
			l++
		}

		// Search right-hand side.
		// Don't allow rhs to move more than one pass lhs.
		for l <= r && nums[r] > pivot {
			r--
		}

		// Still within bounds?
		if l <= r {
			// Swap about pivot-value.
			nums[l], nums[r] = nums[r], nums[l]
			// Advance pointers.
			l++
			r--
		}
	}
	// Move pivot-value into place.
	nums[lo], nums[r] = nums[r], nums[lo]
	// Return pivot-index
	return r
}

//-------------------------------------

func main() {

	var nums []int
	var start time.Time
	var elapsed1, elapsed2 time.Duration
	var elapsed3, elapsed4 time.Duration

	nums = generateSlice(19)
	fmt.Println("Lomuto:")
	fmt.Println("--- Unsorted --- \n", nums)
	start = time.Now()
	quickSortLomuto(nums)
	elapsed1 = time.Since(start)
	fmt.Println("--- Sorted --- \n", nums)
	fmt.Printf("Execution Time : %v\n", elapsed1)
	fmt.Println()

	//---

	nums = generateSlice(19)
	fmt.Println("Hoare1:")
	fmt.Println("--- Unsorted --- \n", nums)
	start = time.Now()
	quickSortHoare1(nums, 0, len(nums)-1)
	elapsed2 = time.Since(start)
	fmt.Println("--- Sorted --- \n", nums)
	fmt.Printf("Execution Time : %v\n", elapsed2)
	fmt.Println()

	fmt.Printf("Ratio = %.4v\n", elapsed1.Seconds()/elapsed2.Seconds())
	fmt.Println()

	//---

	nums = generateSlice(19)
	fmt.Println("Hoare2:")
	fmt.Println("--- Unsorted --- \n", nums)
	start = time.Now()
	quickSortHoare2(nums, 0, len(nums)-1)
	elapsed3 = time.Since(start)
	fmt.Println("--- Sorted --- \n", nums)
	fmt.Printf("Execution Time : %v\n", elapsed3)
	fmt.Println()

	fmt.Printf("Ratio = %.4v\n", elapsed1.Seconds()/elapsed3.Seconds())
	fmt.Println()

	//---

	nums = generateSlice(19)
	fmt.Println("My Hoare:")
	fmt.Println("--- Unsorted --- \n", nums)
	start = time.Now()
	myQuickSort(nums, 0, len(nums)-1)
	elapsed4 = time.Since(start)
	fmt.Println("--- Sorted --- \n", nums)
	fmt.Printf("Execution Time : %v\n", elapsed4)
	fmt.Println()

	fmt.Printf("Ratio = %.4v\n", elapsed1.Seconds()/elapsed4.Seconds())
	fmt.Println()

	//*-*

	/* nums = generateSlice(9)
	fmt.Println("Recursive:")
	fmt.Println("--- Unsorted --- \n", nums)
	start = time.Now()
	quickSort(nums)
	elapsed = time.Since(start)
	fmt.Println("--- Sorted ---\n", nums)
	fmt.Printf("Execution Time : %v\n", elapsed)
	fmt.Println() */

	//*-*

}

//*-*

// generate a slice of size, size filled with random numbers
func generateSlice(size int) []int {
	nums := make([]int, size)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {
		// for negative numbers
		nums[i] = r.Intn(99) - r.Intn(99)
	}
	return nums
}

//*-*
