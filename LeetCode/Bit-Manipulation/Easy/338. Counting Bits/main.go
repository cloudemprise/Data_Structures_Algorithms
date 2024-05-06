package main

import (
	"fmt"
)

// > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > >

/*
338. Counting Bits
Easy

Given an integer n, return an array ans of length n + 1 such that for each i
(0 <= i <= n), ans[i] is the number of 1's in the binary representation of i.

Example 1:

Input: n = 2
Output: [0,1,1]
Explanation:
0 --> 0
1 --> 1
2 --> 10

Example 2:

Input: n = 5
Output: [0,1,1,2,1,2]
Explanation:
0 --> 0
1 --> 1
2 --> 10
3 --> 11
4 --> 100
5 --> 101


Constraints:

    0 <= n <= 105

Follow up:

    It is very easy to come up with a solution with a runtime of O(n log n).
	Can you do it in linear time O(n) and possibly in a single pass?
    Can you do it without using any built-in function
	(i.e., like __builtin_popcount in C++)?
*/

// < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < <

/* // Recursive Solution: ????
func countBits(n int) []int {
	ans := make([]int, n+1)
	fill(ans, 1, 1, n)
	return ans
}

func fill(ans []int, bits, current, n int) {
	if current > n {
		return
	}

	ans[current] = bits
	fill(ans, bits, current<<1, n)
	fill(ans, bits+1, (current<<1)|1, n)
} */

//-------------------------------------

/*
// Solution 3:

Intuitive Understanding:
------------------------

0 -->         0 = 0 =  0/2 + 0 = 0<<1 + 0 = 0 + 0
1 -->         1 = 1 =  1/2 + 1 = 0<<1 + 1 = 0 + 1
2 -->        10 = 1 =  2/2 + 0 = 1<<1 + 0 = 1 + 0
3 -->        11 = 2 =  3/2 + 1 = 1<<1 + 1 = 1 + 1
4 -->       100 = 1 =  4/2 + 0 = 2<<1 + 0 = 1 + 0
5 -->       101 = 2 =  5/2 + 1 = 2<<1 + 1 = 1 + 1
6 -->       110 = 2 =  6/2 + 0 = 3<<1 + 0 = 2 + 0
7 -->       111 = 3 =  7/2 + 1 = 3<<1 + 1 = 2 + 1
8 -->      1000 = 1 =  8/2 + 0 = 4<<1 + 0 = 1 + 0
9 -->      1001 = 2 =  9/2 + 1 = 4<<1 + 1 = 1 + 1
10 ->      1010 = 2 = 10/2 + 0 = 5<<1 + 0 = 2 + 0
*/

/* func countBits(n int) []int {
	result := make([]int, n+1)
	for i := 0; i <= n; i++ {
		result[i] = result[i>>1] + i&1
	}
	return result
} */

/* func countBits(n int) []int {
	result := make([]int, n+1)
	for i := 0; i <= n; i++ {
		result[i] = result[i/2] + i&1
	}
	return result
} */

//-------------------------------------

/* // Solution 2: stdlib bits.OnesCount()
func countBits(n int) []int {
	result := make([]int, n+1)
	for i := 0; i <= n; i++ {
		result[i] = bits.OnesCount(uint(i))
	}
	return result
} */

//-------------------------------------

// Solution 1b: Small Optimisation.
func countBits(n int) []int {
	result := make([]int, n+1)
	for i := 0; i <= n; i++ {
		for j := i; j > 0; j >>= 1 {
			result[i] += j & 1
		}
	}
	return result
}

//-------------------------------------

/* // Solution 1a:
func countBits(n int) []int {
	var result []int
	for i := 0; i <= n; i++ {
		var tmp int
		for j := i; j > 0; j >>= 1 {
			tmp += j & 1
		}
		result = append(result, tmp)
	}
	return result
} */

//-------------------------------------

func main() {

	var num int

	num = 8
	fmt.Printf("Input Number = %d\n", num)
	fmt.Println("Total Number of Ones = ", countBits(num))
	fmt.Println()

	num = 16
	fmt.Printf("Input Number = %d\n", num)
	fmt.Println("Total Number of Ones = ", countBits(num))
	fmt.Println()

}

//*-*
