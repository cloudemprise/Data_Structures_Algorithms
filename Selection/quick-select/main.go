package main

import (
	"fmt"
	"math/rand"
	"time"
)

// > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > >

/*
Quick Select Algorithm uses the same overall approach as quicksort, choosing
one element as a pivot and partitioning the data in two based on the pivot,
accordingly as less than or greater than the pivot. However, instead of
recursing into both sides, as in quicksort, quickselect only recurses into one
side – the side with the element it is searching for. This reduces the average
complexity from O(nlog n) to O(n), with a worst case of O(n^2).
*/

// < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < <

// Tail Recursive Method, finds kth smallest number.
func quickselectLumoto(nums []int, lo, hi, k int) int {
	if lo == hi {
		return nums[lo] // base case
	} else {
		pIndex := partitionLumoto(nums, lo, hi)

		if pIndex < (k - 1) {
			return quickselectLumoto(nums, pIndex+1, hi, k)
		} else if pIndex > (k - 1) {
			return quickselectLumoto(nums, lo, pIndex-1, k)
		} else {
			return nums[pIndex]
		}
	}
}

//*-*

// Iterative Lumoto Partiton Algo.
func partitionLumoto(nums []int, lo, hi int) int {

	// Select last element as pivot value.
	pivot := nums[hi]

	// Search for & move values less than pivot-value down.
	for i := lo; i < hi; i++ {
		if nums[i] < pivot {
			nums[lo], nums[i] = nums[i], nums[lo]
			lo++ // move ptr forward after swap.
		}
	}
	// Move pivot-value into its proper position.
	nums[lo], nums[hi] = nums[hi], nums[lo]
	return lo
}

//-------------------------------------

// Tail Recursive Method, finds kth smallest number.
func quickselectHoare(nums []int, lo, hi, k int) int {
	if lo == hi {
		return nums[lo] // base case
	} else {
		pIndex := partitionHoare(nums, lo, hi)

		if pIndex < (k - 1) {
			return quickselectHoare(nums, pIndex+1, hi, k)
		} else if pIndex > (k - 1) {
			return quickselectHoare(nums, lo, pIndex-1, k)
		} else {
			return nums[pIndex]
		}
	}
}

//*-*

// Iterative Hoare Partition Algo.
func partitionHoare(nums []int, lo, hi int) int {

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

	var nums1 []int
	var kth int
	var start time.Time
	var elapsed time.Duration

	n := 11
	k := n / 2
	nums1 = generateSliceUnique(n)
	nums2 := make([]int, len(nums1))
	copy(nums2, nums1)

	//--

	fmt.Println("Lumoto: ")
	fmt.Printf("k = %v\n", k)
	fmt.Println("--- Unsorted --- \n", nums1)
	start = time.Now()
	kth = quickselectLumoto(nums1, 0, len(nums1)-1, k)
	elapsed = time.Since(start)
	fmt.Println("--- Selected --- \n", nums1)
	fmt.Printf("Execution Time : %v\n", elapsed)
	fmt.Printf("k-th element value = %v\n", kth)
	fmt.Println()

	//*-*

	//nums2 = generateSliceUnique(9)
	fmt.Println("Hoare: ")
	fmt.Printf("k = %v\n", k)
	fmt.Println("--- Unsorted --- \n", nums2)
	start = time.Now()
	kth = quickselectHoare(nums2, 0, len(nums1)-1, k)
	elapsed = time.Since(start)
	fmt.Println("--- Selected --- \n", nums2)
	fmt.Printf("Execution Time : %v\n", elapsed)
	fmt.Printf("k-th element value = %v\n", kth)
	fmt.Println()

	//*-*

}

//*-*

// Generate a slice of size, size filled with random numbers
func generateSlice(size int) []int {
	nums := make([]int, size)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {
		// for negative numbers
		nums[i] = r.Intn(99) - r.Intn(99)
		//nums[i] = r.Intn(99)
	}
	//nums = r.Perm(9)
	return nums
}

// Generate a slice of size random numbers w/o repeats.
func generateSliceUnique(size int) []int {
	nums := make([]int, size)
	generated := make(map[int]bool)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; {
		num := r.Intn(99)
		if !generated[num] {
			generated[num] = true
			nums[i] = num
			i++
		}
	}
	return nums
}

//*-*
