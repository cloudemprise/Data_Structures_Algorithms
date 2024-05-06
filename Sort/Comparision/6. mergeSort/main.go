package main

import (
	"fmt"
	"math/rand"
	"time"
)

// > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > >

/*
Merge sort Algorithm

*/

// < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < <

func mergeSort(nums []int) []int {
	ln := len(nums)
	if ln <= 1 {
		return nums
	}
	lhs := mergeSort(nums[:ln/2])
	rhs := mergeSort(nums[ln/2:])
	return merge(lhs, rhs)
}

//-------------------------------------

/* // Recursive Version - Rubio-Sanchez Pg. 177
func merge(lhs, rhs []int) []int {
	if len(lhs) == 0 {
		return rhs
	}
	if len(rhs) == 0 {
		return lhs
	}

	if lhs[0] < rhs[0] {
		result := append([]int{lhs[0]}, merge(lhs[1:], rhs)...)
		return result
	} else {
		result := append([]int{rhs[0]}, merge(lhs, rhs[1:])...)
		return result
	}
} */

// My modified version.
func merge(lhs, rhs []int) []int {

	// Preallocate capacity
	result := make([]int, 0, len(lhs)+len(rhs))

	// Merge in sorted order.
	for len(lhs) > 0 && len(rhs) > 0 {
		if lhs[0] < rhs[0] {
			result = append(result, lhs[0])
			lhs = lhs[1:]
		} else {
			result = append(result, rhs[0])
			rhs = rhs[1:]
		}
	}

	// Append remaining values.
	if len(lhs) > 0 {
		result = append(result, lhs...)
	}
	if len(rhs) > 0 {
		result = append(result, rhs...)
	}
	return result
}

/* // Karumanchi Pg.321
func merge(lhs, rhs []int) []int {
	lnL := len(lhs)
	lnR := len(rhs)

	// Pre-allocate mem
	result := make([]int, lnL+lnR)

	// result index approach.
	for i := 0; i < lnL+lnR; i++ {

		if len(lhs) > 0 && len(rhs) > 0 {
			// Merge in sorted order.
			switch {
			case lhs[0] < rhs[0]:
				result[i] = lhs[0]
				lhs = lhs[1:]
			case lhs[0] >= rhs[0]:
				result[i] = rhs[0]
				rhs = rhs[1:]
			}

		} else {
			// Append remaining values.
			if len(lhs) > 0 {
				result[i] = lhs[0]
				lhs = lhs[1:]
			}
			if len(rhs) > 0 {
				result[i] = rhs[0]
				rhs = rhs[1:]
			}
		}
	}
	return result
} */

/* // Iterative Version golangprograms.com
func merge(lhs, rhs []int) []int {
	result := make([]int, len(lhs)+len(rhs))

	i := 0
	for len(lhs) > 0 && len(rhs) > 0 {
		if lhs[0] < rhs[0] {
			result[i] = lhs[0]
			lhs = lhs[1:]
		} else {
			result[i] = rhs[0]
			rhs = rhs[1:]
		}
		i++
	}

	for j := 0; j < len(lhs); j++ {
		result[i] = lhs[j]
		i++
	}
	for j := 0; j < len(rhs); j++ {
		result[i] = rhs[j]
		i++
	}

	return result
} */

//-------------------------------------

func main() {

	var nums []int
	var start time.Time
	var elapsed1 time.Duration
	//var elapsed3, elapsed4 time.Duration

	nums = generateSlice(9)
	fmt.Println("Merge Sort:")
	fmt.Println("--- Unsorted --- \n", nums)
	start = time.Now()
	nums = mergeSort(nums)
	elapsed1 = time.Since(start)
	fmt.Println("--- Sorted --- \n", nums)
	fmt.Printf("Execution Time : %v\n", elapsed1)
	fmt.Println()

	//---

	/* nums = generateSlice(19)
	fmt.Println("Hoare1:")
	fmt.Println("--- Unsorted --- \n", nums)
	start = time.Now()
	quickSortHoare1(nums, 0, len(nums)-1)
	elapsed2 = time.Since(start)
	fmt.Println("--- Sorted --- \n", nums)
	fmt.Printf("Execution Time : %v\n", elapsed2)
	fmt.Println()

	fmt.Printf("Ratio = %.4v\n", elapsed1.Seconds()/elapsed2.Seconds())
	fmt.Println() */

	//---

	/* nums = generateSlice(19)
	fmt.Println("Hoare2:")
	fmt.Println("--- Unsorted --- \n", nums)
	start = time.Now()
	quickSortHoare2(nums, 0, len(nums)-1)
	elapsed3 = time.Since(start)
	fmt.Println("--- Sorted --- \n", nums)
	fmt.Printf("Execution Time : %v\n", elapsed3)
	fmt.Println()

	fmt.Printf("Ratio = %.4v\n", elapsed1.Seconds()/elapsed3.Seconds())
	fmt.Println() */

	//---

	/* nums = generateSlice(19)
	fmt.Println("My Hoare:")
	fmt.Println("--- Unsorted --- \n", nums)
	start = time.Now()
	myQuickSort(nums, 0, len(nums)-1)
	elapsed4 = time.Since(start)
	fmt.Println("--- Sorted --- \n", nums)
	fmt.Printf("Execution Time : %v\n", elapsed4)
	fmt.Println()

	fmt.Printf("Ratio = %.4v\n", elapsed1.Seconds()/elapsed4.Seconds())
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
