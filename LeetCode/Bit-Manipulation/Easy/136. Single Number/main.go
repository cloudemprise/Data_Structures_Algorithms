package main

import "fmt"

// > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > >

/*
136. Single Number
Easy

Given a non-empty array of integers nums, every element appears twice except
for one. Find that single one.

You must implement a solution with a linear runtime complexity and use only
constant extra space.


Example 1:

Input: nums = [2,2,1]
Output: 1

Example 2:

Input: nums = [4,1,2,1,2]
Output: 4

Example 3:

Input: nums = [1]
Output: 1


Constraints:

    1 <= nums.length <= 3 * 104
    -3 * 104 <= nums[i] <= 3 * 104
    Each element in the array appears twice except for one element which
	appears only once.
*/

/*
Intuitive Understanding:
------------------------

*/

// < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < <

func singleNumber(nums []int) int {
	var answer int
	for _, num := range nums {
		answer ^= num
	}
	return answer
}

//-------------------------------------

func main() {

	var nums []int

	nums = []int{2, 2, 1}
	fmt.Printf("Input = %#v\n", nums)
	fmt.Printf("Answer = %#v", singleNumber(nums))
	fmt.Println()

	nums = []int{4, 1, 2, 1, 2}
	fmt.Printf("Input = %#v\n", nums)
	fmt.Printf("Answer = %#v", singleNumber(nums))
	fmt.Println()
}

//*-*
