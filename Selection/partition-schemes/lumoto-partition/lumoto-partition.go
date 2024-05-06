package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
QUICK SELECT Algorithm uses the same overall approach as quicksort, choosing
one element as a pivot and partitioning the data in two based on the pivot,
accordingly as less than or greater than the pivot. However, instead of
recursing into both sides, as in quicksort, quickselect only recurses into one
side – the side with the element it is searching for. This reduces the average
complexity from O(nlogn) to O(n), with a worst case of O(n^2).
*/

// lomutoPartition partitions an unordered list about a chosen pivot element
// value and then returns the index of that pivot element after moving it into
// it final position.
func lomutoPartition(nums []int, lo, hi int) int {

	// Select last element as pivot value.
	pivot := nums[hi]

	// Search for & move values less than pivot-value down.
	for i := lo; i < hi; i++ {
		if nums[i] < pivot {
			nums[lo], nums[i] = nums[i], nums[lo]
			lo++ // move ptr forward after swap.
		}
	}
	// Move pivot-value into its final position.
	nums[lo], nums[hi] = nums[hi], nums[lo]
	return lo
}

///

// generateRandomNums returns a list of random integers, including negative numbers.
func generateRandomNums(size int) []int {

	// New random number generator.
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	nums := make([]int, size)
	for i := 0; i < size; i++ {
		nums[i] = r.Intn(99) - r.Intn(99) // Includes negative numbers.
	}
	return nums
}

// generateUniqueRandomNums returns a list of unique random integers, including
// negative numbers.
func generateUniqueRandomNums(size int) []int {

	// New random number generator.
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	generated := make(map[int]bool)

	nums := make([]int, size)
	for i := 0; i < size; {
		num := r.Intn(99) - r.Intn(99) // Includes negative numbers.
		if !generated[num] {
			generated[num] = true
			nums[i] = num
			i++
		}
	}
	return nums
}

///

func main() {

	var nums []int
	var pivot int
	var start time.Time
	var elapsed1 time.Duration

	// Special Cases:
	//nums = []int{94, 2, 26, 10, 65, 22, 5, 35, 38, 88, 84}
	//nums = []int{93, 91, 64, 81, 92, 44, 75, 69, 72}
	//nums = []int{8, 56, 94, 7, 45, 61, 25, 81, 49, 22, 96}
	//nums = []int{2, 3, 4, 1}
	//nums = []int{4, 2, 1, 3}
	nums = []int{1, 3, 2, 4}
	//nums = []int{3, 2, 1}
	//nums = []int{1, 2, 3}

	//nums = generateUniqueRandomNums(11)

	fmt.Println("Lomuto Partition Scheme: ")
	fmt.Println("Original array:\t\t", nums)
	start = time.Now()
	pivot = lomutoPartition(nums, 0, len(nums)-1)
	elapsed1 = time.Since(start)
	fmt.Println("After partitioning:\t", nums)
	fmt.Println("Pivot value:\t\t", nums[pivot])
	fmt.Println("Pivot index:\t\t", pivot)
	fmt.Println("Left sub-array:\t\t", nums[:pivot])
	fmt.Println("Right sub-array:\t", nums[pivot+1:])
	fmt.Println("Execution Time:\t\t", elapsed1)
	fmt.Println()

}
