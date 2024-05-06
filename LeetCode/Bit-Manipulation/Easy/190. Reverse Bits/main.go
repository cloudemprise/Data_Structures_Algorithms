package main

import "fmt"

// > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > >

/*
190. Reverse Bits
Easy

Reverse bits of a given 32 bits unsigned integer.

Example 1:

Input:            n = 00000010100101000001111010011100
Output:    964176192 (00111001011110000010100101000000)
Explanation:
The input binary string 00000010100101000001111010011100 represents the unsigned
integer 43261596, so return 964176192 which its binary representation is
00111001011110000010100101000000.

Example 2:

Input:            n = 11111111111111111111111111111101
Output:   3221225471 (10111111111111111111111111111111)
Explanation:
The input binary string 11111111111111111111111111111101 represents the unsigned
integer 4294967293, so return 3221225471 which its binary representation is
10111111111111111111111111111111.

Constraints:

    The input must be a binary string of length 32

Follow up: If this function is called many times, how would you optimize it?
*/

/*
Intuitive Understanding:
------------------------

*/

// < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < <

// Solution 2: Without Bit Operations
func reverseBits(num uint32) uint32 {
	var result uint32
	for i := 0; i < 32; i++ {
		result *= 2
		result += num % 2
		num /= 2
	}
	return result
}

//-------------------------------------

/* // Solution 1: Bit Operations
func reverseBits(num uint32) uint32 {
	var result uint32
	for i := 0; i < 32; i++ {
		result <<= 1   // rhs shift : multiply by 2
		lsb := num & 1 // least significant bit
		result |= lsb  // modify lsb
		num >>= 1      // lhs shift : divide by 2
	}
	return result
} */

//-------------------------------------

func main() {

	var num uint32

	num = 0b01010101
	fmt.Printf("Input Binary = %08b\n", num)
	fmt.Println("--------------------------")
	fmt.Printf("Output Binary = %08b\n", reverseBits(num))
	fmt.Println("--------------------------")
	fmt.Println()

	/* 	num = 0b00000011
	   	fmt.Printf("Input Binary = %032b\n", num)
	   	fmt.Printf("Output Binary = %032b\n", reverseBits(num))
	   	fmt.Println() */

}

//*-*
