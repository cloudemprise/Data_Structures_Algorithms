package main

import "fmt"

/*
401. Binary Watch - Easy

A binary watch has 4 LEDs on the top to represent the hours (0-11), and 6 LEDs
on the bottom to represent the minutes (0-59). Each LED represents a zero or
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

// Solution 1: Backtracking recusive closure using a bitmap to represent LED
// time sequences.
const totalLEDs = 10

func readBinaryWatch(turnedOn int) (results []string) {

	// Accumulate all valid solutions. So as to avoid repeated permutations.
	cache := make(map[int]bool)

	// Implement a backtracking recursive closure.
	var backtrack func(int, map[int]bool, int)
	backtrack = func(soln int, isOn map[int]bool, numOfOnLEDs int) {
		if numOfOnLEDs == turnedOn { // base case.
			// Solution found, so copy into cache.
			cache[soln] = true
			return
		} else {
			// Iterate through every LED index.
			for i := 0; i < totalLEDs; i++ {
				mask := 1 << i // bit-mask

				// Only consider LEDs that are Off.
				if !isOn[i] {
					// Mark this LED as On.
					isOn[i] = true
					// Add LED to partial solution.
					soln ^= mask

					// Check if it is indeed a valid time sequence.
					hrs, mins := calcHrsMins(soln)
					if hrs < 12 && mins < 60 {

						// Check if time sequence is unique.
						if !cache[soln] {
							// Search for other solutions.
							backtrack(soln, isOn, numOfOnLEDs+1)
						}
					}

					// Now turn this LED Off.
					isOn[i] = false
					// Remove LED from partial solution.
					soln &^= mask
				}
			}
		}
	}

	isOn := make(map[int]bool)             // Remember which LEDs are On/Off.
	solution := 0                          // bitmap of time permutations.
	numOfOnLEDs := 0                       // No. of LEDs currently turned ON.
	backtrack(solution, isOn, numOfOnLEDs) // Call recursive closure.

	// Create output string from bitmap cache.
	for soln := range cache {
		hrs, mins := calcHrsMins(soln)
		result := fmt.Sprintf("%d:%02d", hrs, mins)
		results = append(results, result)
	}
	return results
}

// calcHrsMins converts the solution bitmap into a time statement.
func calcHrsMins(soln int) (hrs, mins int) {
	// 10-bit summation encoder.
	for i := 0; i < totalLEDs; i++ {
		switch soln & (1 << i) {
		case 1: // 0b0000000001
			mins += 1
		case 2: // 0b0000000010
			mins += 2
		case 4: // 0b0000000100
			mins += 4
		case 8: // 0b0000001000
			mins += 8
		case 16: // 0b0000010000
			mins += 16
		case 32: // 0b0000100000
			mins += 32
		case 64: // 0b0001000000
			hrs += 1
		case 128: // 0b0010000000
			hrs += 2
		case 256: // 0b0100000000
			hrs += 4
		case 512: // 0b1000000000
			hrs += 8
		}
	}
	return hrs, mins
}

//-------------------------------------

func main() {

	fmt.Println(readBinaryWatch(1))

}

//*-*
