package main

import (
	"fmt"
	"math/rand"
	"time"
)

// > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > >

/*
268. Missing Number
Easy

Given an array nums containing n distinct numbers in the range [0, n], return
the only number in the range that is missing from the array.


Example 1:

Input: nums = [3,0,1]
Output: 2
Explanation: n = 3 since there are 3 numbers, so all numbers are in the range
[0,3]. 2 is the missing number in the range since it does not appear in nums.

Example 2:

Input: nums = [0,1]
Output: 2
Explanation: n = 2 since there are 2 numbers, so all numbers are in the range
[0,2]. 2 is the missing number in the range since it does not appear in nums.

Example 3:

Input: nums = [9,6,4,2,3,5,7,0,1]
Output: 8
Explanation: n = 9 since there are 9 numbers, so all numbers are in the range
[0,9]. 8 is the missing number in the range since it does not appear in nums.

Constraints:

    n == nums.length
    1 <= n <= 104
    0 <= nums[i] <= n
    All the numbers of nums are unique.

Follow up: Could you implement a solution using only O(1) extra space complexity
and O(n) runtime complexity?
*/

// < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < <

/*
Explanation:
1^1=0
1^2^1^2=0
1^2^3^1^2^3=0
...etc.
Therefore eg. 1^2^3^1^3=2
*/

/* // Solution 3:
// Time Complexity: O(n)
// Space Complexity: O(1)
func missingNumber(nums []int) int {
	var result int
	for i := 0; i < len(nums); i++ {
		result ^= nums[i] ^ (i + 1)
	}
	return result
} */

//-------------------------------------

// Solution 2:
// Time Complexity: O(n)
// Space Complexity: O(1)
func missingNumber(nums []int) int {
	ln := len(nums)
	var want, got int
	for i := 0; i < ln; i++ {
		got += nums[i]
		want += i
	}
	// add an extra loop iteration.
	return (want + ln) - got
}

//-------------------------------------

/* // Solution 1:
// Time Complexity: O(n^2)
// Space Complexity : O(n)
func missingNumber(nums []int) int {
	// populate a map
	cache := make(map[int]struct{})
	for _, num := range nums {
		cache[num] = struct{}{}
	}

	// search for missing consecutive number in map
	for i := 0; i <= len(nums); i++ {
		if _, ok := cache[i]; !ok {
			return i
		}
	}
	return -1 // unknown error
} */

//-------------------------------------

func main() {

	// generate a source of random numbers.
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// select a random number within the interval [1,x)
	n := r.Intn(9) + 1
	// generate a slice of n ints within the range.
	nums := r.Perm(n)
	// remove one element to satisfy problem space.
	nums = nums[:len(nums)-1]

	fmt.Println("The input elements = ", nums)
	fmt.Println("The missing element = ", missingNumber(nums))

}

//*-*
