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

// Solution 2: Backtracking recusive closure using a bitmap to represent LED
// time sequences (simplified).
const totalLEDs = 10

func readBinaryWatch(turnedOn int) (results []string) {

	// Accumulate all valid solutions. So as to avoid repeated permutations.
	cache := make(map[int]bool)

	var backtrack func(int, int)
	backtrack = func(soln int, numOfOnLEDs int) {
		if numOfOnLEDs == turnedOn { // base case.
			// Solution found, so copy into cache.
			cache[soln] = true
			return
		}

		// Iterate through every LED permutation.
		for mask := 1; mask < (1 << totalLEDs); mask <<= 1 {

			// Ignore LEDs that are On.
			if !(soln&mask > 0) {
				// Add LED to partial solution.
				soln ^= mask

				// Check if it is indeed a valid time sequence.
				hrs, mins := calcHrsMins(soln)
				if hrs < 12 && mins < 60 {

					// Check if time sequence is unique.
					if !cache[soln] {
						// Search for other solutions.
						backtrack(soln, numOfOnLEDs+1)
					}
				}

				// Remove LED from partial solution.
				soln &^= mask
			}
		}
	}

	solution := 0                    // bitmap of time permutations.
	numOfOnLEDs := 0                 // No. of LEDs currently turned ON.
	backtrack(solution, numOfOnLEDs) // Call recursive closure.

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
