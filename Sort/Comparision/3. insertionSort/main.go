package main

import (
	"fmt"
	"math/rand"
	"time"
)

// > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > >

/*
Insertion Sort Algorithm
time complexity = O(n^2)
space complexity = O(1)

This algorithm assumes that the left portion of the list has already been
sorted, while the right portion remains unsorted. We itterates through the
unsorted portion of the list, picking one element at a time. With this
element, we go through the sorted portion of the list and insert it in the
right order so that the sorted portion of the list remains sorted.
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

/* // Iterative Insertion Sort Algorithm
func insertionSort(slice []int) {
	ln := len(slice)
	for i := 1; i <= ln-1; i++ {
		j := i
		for j > 0 {
			if slice[j-1] > slice[j] {
				slice[j-1], slice[j] = slice[j], slice[j-1]
			}
			j--
		}
	}
} */
// Iterative Insertion Sort Algorithm
func insertionSort(slice []int) {
	ln := len(slice)
	for i := 1; i <= ln-1; i++ {
		for j := i; j > 0; j-- {
			if slice[j-1] > slice[j] {
				slice[j-1], slice[j] = slice[j], slice[j-1]
			}
		}
	}
}

//*-*

func main() {

	slice := generateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	insertionSort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice)
}
