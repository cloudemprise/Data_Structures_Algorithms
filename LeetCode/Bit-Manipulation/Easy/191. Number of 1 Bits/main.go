package main

import "fmt"

// > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > >

/*
191. Number of 1 Bits
Easy

Write a function that takes the binary representation of an unsigned integer
and returns the number of '1' bits it has (also known as the Hamming weight).

Example 1:

Input: n = 00000000000000000000000000001011
Output: 3
Explanation: The input binary string 00000000000000000000000000001011 has a
total of three '1' bits.

Example 2:

Input: n = 00000000000000000000000010000000
Output: 1
Explanation: The input binary string 00000000000000000000000010000000 has a
total of one '1' bit.

Example 3:

Input: n = 11111111111111111111111111111101
Output: 31
Explanation: The input binary string 11111111111111111111111111111101 has a
total of thirty one '1' bits.

Constraints:

    The input must be a binary string of length 32.
*/

/*
Intuitive Understanding:
------------------------

*/

// < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < <

// Solution 3:
func hammingWeight(num uint32) int {
	var total uint32
	for num != 0 {
		total += num & 1
		num >>= 1
	}
	return int(total)
}

//-------------------------------------

/* // Solution 2:
func hammingWeight(num uint32) int {
	var total int
	for num > 0 {
		total += int(num & 1)
		num = num >> 1
	}
	return total
} */

//-------------------------------------

/* // Solution 1:
func hammingWeight(num uint32) int {
	var total int
	for i := 0; i < 32; i++ {
		total += int(num & 1)
		num = num >> 1
	}
	return total
} */

//-------------------------------------

func main() {

	var num uint32

	num = 0b00001101
	fmt.Printf("Input Binary = %08b\n", num)
	fmt.Println("Total Number of Ones = ", hammingWeight(num))
	fmt.Println()

	num = 0b10101010
	fmt.Printf("Input Binary = %08b\n", num)
	fmt.Println("Total Number of Ones = ", hammingWeight(num))
	fmt.Println()

}

//*-*
