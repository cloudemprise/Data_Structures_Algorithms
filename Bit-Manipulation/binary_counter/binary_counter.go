package main

import (
	"fmt"
)

// > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > >

/*
Comments
--------

*/

// < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < <

//-------------------------------------

// Solution 2: Iterative subsets
func binaryCounter2(numOfBits int) [][]int {

	var subsets [][]int // accumulator.

	// Total number of subsets = 2^x
	// A binary (integer) power function.
	numOfSubsets := 1 << numOfBits
	// Big-endian binary counter
	for i := 0; i < numOfSubsets; i++ {
		var batch []int // local accumulator.
		for j := 0; j < numOfBits; j++ {
			mask := 1 << j // bit mask
			// test bit
			if (i & mask) > 0 {
				batch = append(batch, 1)
			} else {
				batch = append(batch, 0)
			}
		}
		subsets = append(subsets, batch)
	}
	return subsets
}

// Solution 1: Iterative subsets
func binaryCounter1(numOfBits int) [][]int {

	var subsets [][]int // accumulator.

	// Total number of subsets = 2^x
	// A binary (integer) power function.
	numOfSubsets := 1 << numOfBits
	// Big-endian binary counter
	for i := 0; i < numOfSubsets; i++ {
		value := i
		var batch []int // local accumulator.
		for j := 0; j < numOfBits; j++ {
			lsb := value & 1
			batch = append(batch, lsb)
			value >>= 1
		}
		subsets = append(subsets, batch)
	}
	return subsets
}

func main() {

	fmt.Println(binaryCounter1(3))
	fmt.Println()

	//*-*
	fmt.Println(binaryCounter2(3))
	fmt.Println()

	//*-*

}

//*-*
