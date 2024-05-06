package main

import (
	"fmt"
	"math/rand"
	"time"
)

// > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > >
/*
Bubble Sort Algorithm
time complexity = O(n^2)
space complexity = O(1)

An algorithm that repeatedly steps through a list and compares each pair of
adjacent elements and swaps them if they are in the wrong order.

time complexity = O(n^2)
space complexity = O(1)

Each pass through the list will place the next largest element value in its
proper place, thus: for i := 1; ln > i; i++ {...}
*/
// < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < <

// generate a slice of size, size filled with random numbers
func generateSlice(size int) []int {
	slice := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		// for negative numbers
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

//*-*

// Iterative Bubble Sort Algorithm
func bubbleSort(slice []int) {
	ln := len(slice)
	var sorted bool
	for !sorted {
		sorted = true
		for i := 1; i < ln; i++ {
			if slice[i-1] > slice[i] {
				slice[i-1], slice[i] = slice[i], slice[i-1]
				sorted = false
			}
		}
	}
}

//*-*

func main() {

	slice := generateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	bubbleSort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice)
}
