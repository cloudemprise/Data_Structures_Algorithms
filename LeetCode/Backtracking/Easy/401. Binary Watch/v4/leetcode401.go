package main

import (
	"fmt"
	"math/bits"
)

/*
401. Binary Watch - Easy

A binary watch has 4 LEDs on the top to represent the hours (0-11), and 6 LEDs
on the bottom to represent the minutes (0-59). Each LEDindex represents a zero or
one, with the least significant bit on the right.

For example, the below binary watch reads "4:51".

Given an integer turnedOn which represents the number of LEDs that are
currently on (ignoring the PM), return all possible times the watch could
represent. You may return the answer in any order.

The hour must not contain a leading zero.

    For example, "01:00" is not valid. It should be "1:00".

The minute must be consist of two digits and may contain a leading zero.

    For example, "10:2" is not valid. It should be "10:02".


Example 1:

Input: turnedOn = 1
Output: ["0:01","0:02","0:04","0:08","0:16","0:32","1:00","2:00","4:00","8:00"]

Example 2:

Input: turnedOn = 9
Output: []


Constraints:

    0 <= turnedOn <= 10
*/

/*
//-------------------------------------
// Leetcode reference solution
func readBinaryWatch(turnedOn int) []string {
    var times []string
    for h := 0; h < 12; h++ {
        i := bits.OnesCount(uint(h))
        for m := 0; m < 60; m++ {
            j := bits.OnesCount(uint(m))
            if turnedOn == i + j {
               times = append(times, fmt.Sprintf("%d:%02d", h, m))
            }
        }
    }
    return times
}
//-------------------------------------
*/

/*
Bitmap
LEDindex : 0b0000000001 :  1 mins => index=0
LEDindex : 0b0000000010 :  2 mins => index=1
LEDindex : 0b0000000100 :  4 mins => index=2
LEDindex : 0b0000001000 :  8 mins => index=3
LEDindex : 0b0000010000 : 16 mins => index=4
LEDindex : 0b0000100000 : 32 mins => index=5
LEDindex : 0b0001000000 :  1 hrs  => index=6
LEDindex : 0b0010000000 :  2 hrs  => index=7
LEDindex : 0b0100000000 :  4 hrs  => index=8
LEDindex : 0b1000000000 :  8 hrs  => index=9
*/

// Solution 4: Iterative Brute force.
func readBinaryWatch(turnedOn int) (results []string) {

	for hrs := 0; hrs < 12; hrs++ {
		hrsBits := bits.OnesCount(uint(hrs))

		for mins := 0; mins < 60; mins++ {
			minsBits := bits.OnesCount(uint(mins))

			// Only consider totals that satisfy the constraint.
			if turnedOn == (hrsBits + minsBits) {

				results = append(results, fmt.Sprintf("%d:%02d", hrs, mins))
			}
		}
	}
	return results
}

//-------------------------------------

func main() {

	fmt.Println(readBinaryWatch(1))

}

//*-*
