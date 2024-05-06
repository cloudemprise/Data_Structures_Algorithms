package main

import "fmt"

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
func readBinaryWatch(num int) []string {
	var res []string
	doRead(&res, num, 0, 0, 1)
	return res
}

func doRead(res *[]string, num int, hour , minute int, pos uint){

	if num == 0{
		*res = append(*res, strconv.Itoa(hour) + ":" + fmt.Sprintf("%02d", minute))
		return
	}


	for i:=pos;i<=10;i++{
		if i >=5 {
			if minute + 1<<(i-5) >=60{
				return
			}
			doRead(res, num-1, hour, minute+ 1<<(i-5), i + 1)
		}else {
			if hour + 1<<(i-1) >= 12{
				continue
			}
			doRead(res, num-1, hour + 1<<(i-1), minute, i + 1)
		}
	}
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

const totalLEDs = 10

// Solution 3a: Backtracking recusive closure with separate hrs and mins variables.
func readBinaryWatch(turnedOn int) (results []string) {

	// Implement a backtracking recursive closure.
	var backtrack func(int, int, int, int)
	backtrack = func(hrs, mins, numOfOnLEDs, skipDuplicates int) {

		if numOfOnLEDs == turnedOn { // base case.
			// Solution found, so copy into results string.
			results = append(results, fmt.Sprintf("%d:%02d", hrs, mins))
			return

		} else {

			// Iterate through permutations: LEDindex 0 -> 9
			for LEDindex := skipDuplicates; LEDindex < totalLEDs; LEDindex++ {

				if LEDindex >= 6 {

					// Hours:
					// LEDindex : 0b0001000000 :  1 hrs  => index=6
					// LEDindex : 0b0010000000 :  2 hrs  => index=7
					// LEDindex : 0b0100000000 :  4 hrs  => index=8
					// LEDindex : 0b1000000000 :  8 hrs  => index=9

					// Add hrs(LEDindex) to partial solution.
					hrs += (1 << (LEDindex - 6))
					if hrs < 12 { // Ignore invalid solutions.
						backtrack(hrs, mins, numOfOnLEDs+1, LEDindex+1)
					}
					// Remove hrs(LEDindex) from partial solution.
					hrs -= (1 << (LEDindex - 6))

				} else {

					// Minutes:
					// LEDindex : 0b0000000001 :  1 mins => index=0
					// LEDindex : 0b0000000010 :  2 mins => index=1
					// LEDindex : 0b0000000100 :  4 mins => index=2
					// LEDindex : 0b0000001000 :  8 mins => index=3
					// LEDindex : 0b0000010000 : 16 mins => index=4
					// LEDindex : 0b0000100000 : 32 mins => index=5

					// Add mins(LEDindex) to partial solution.
					mins += (1 << LEDindex)
					if mins < 60 { // Ignore invalid solutions.
						backtrack(hrs, mins, numOfOnLEDs+1, LEDindex+1)
					}
					// Remove mins(LEDindex) from partial solution.
					mins -= (1 << LEDindex)
				}
			}
		}
	}

	numOfOnLEDs := 0    // Number of currently turned On LEDs.
	skipDuplicates := 0 // Window slider to ignore duplicate permutations.
	// Call recursive closure.
	backtrack(0, 0, numOfOnLEDs, skipDuplicates)

	return results
}

//-------------------------------------

func main() {

	fmt.Println(readBinaryWatch(1))

}

//*-*
