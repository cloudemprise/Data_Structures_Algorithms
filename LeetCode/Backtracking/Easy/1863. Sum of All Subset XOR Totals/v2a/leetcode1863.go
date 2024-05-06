package main

import (
	"fmt"
)

/*
The XOR total of an array is defined as the bitwise XOR of all its elements,
or 0 if the array is empty.

For example, the XOR total of the array [2,5,6] is 2 XOR 5 XOR 6 = 1.

Given an array nums, return the sum of all XOR totals for every subset of nums.

Note: Subsets with the same elements should be counted multiple times.

An array a is a subset of an array b if a can be obtained from b by deleting
some (possibly zero) elements of b.


Example 1:

Input: nums = [1,3]
Output: 6
Explanation: The 4 subsets of [1,3] are:

- The empty subset has an XOR total of 0.
- [1] has an XOR total of 1.
- [3] has an XOR total of 3.
- [1,3] has an XOR total of 1 XOR 3 = 2.

0 + 1 + 3 + 2 = 6


Example 2:

Input: nums = [5,1,6]
Output: 28
Explanation: The 8 subsets of [5,1,6] are:

- The empty subset has an XOR total of 0.
- [5] has an XOR total of 5.
- [1] has an XOR total of 1.
- [6] has an XOR total of 6.
- [5,1] has an XOR total of 5 XOR 1 = 4.
- [5,6] has an XOR total of 5 XOR 6 = 3.
- [1,6] has an XOR total of 1 XOR 6 = 7.
- [5,1,6] has an XOR total of 5 XOR 1 XOR 6 = 2.

0 + 5 + 1 + 6 + 4 + 3 + 7 + 2 = 28

Example 3:

Input: nums = [3,4,5,6,7,8]
Output: 480
Explanation: The sum of all XOR totals for every subset is 480.


Constraints:

    1 <= nums.length <= 12
    1 <= nums[i] <= 20
*/

/*
Induction: nums = int{1, 3, 5}

             Inclusion Tree                          Exclusion Tree

                                     {1,3,5}
                    {1}                                     {}
        {1,3}                 {1}                 {3}                {}
{1,3,5}      {1,3}      {1,5}      {1}      {3,5}      {3}      {5}      {}

*/

// Solution 2a. Iterative subsets.
func subsetXORSum(nums []int) (sum int) {

	// The total number of subsets, i.e. 2^x
	// A binary (integer) power function.
	numOfSubsets := 1 << len(nums)
	// Big-endian binary counter
	for i := 0; i < numOfSubsets; i++ {
		val := i
		// Create a subset from val.
		var subset []int
		for idx := 0; idx < len(nums); idx++ {
			lsb := val & 1
			subset = append(subset, nums[idx]*lsb)
			val >>= 1
		}
		// Process the subsets.
		var xOR int // local accumulator.
		for _, val := range subset {
			// Business logic.
			xOR ^= val
		}
		// Accumulate subsets xOR.
		sum += xOR
	}
	return sum
}

func main() {

	var input []int

	input = []int{1, 3}
	fmt.Printf("Input=%v\n", input)
	fmt.Printf("Want: 6 \t Got: %v\n", subsetXORSum(input))
	fmt.Println()

	//*-*

	input = []int{5, 1, 6}
	fmt.Printf("Input=%v\n", input)
	fmt.Printf("Want: 28 \t Got: %v\n", subsetXORSum(input))
	fmt.Println()

	//*-*

	input = []int{3, 4, 5, 6, 7, 8}
	fmt.Printf("Input=%v\n", input)
	fmt.Printf("Want: 480 \t Got: %v\n", subsetXORSum(input))
	fmt.Println()

}