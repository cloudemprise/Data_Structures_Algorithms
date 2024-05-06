package main

import (
	"fmt"
)

/*
Logic Breakdown:
------------------------

- We initialize an empty stack and preallocate a waitDays array of the same length
as the input temperatures.
- We iterate through the temperatures array from left to right. For each
temperature, we check if it is greater than the temperature at the index on top
of the stack.
- If it is, we update the result array at the index on top of the stack with the
difference between the current index and the index on the top of the stack
(indicating the number of days until a warmer temperature is encountered).
- We continue this process until the stack is empty or the current temperature
is not greater than the temperature on top of the stack.
- Finally, we push the current index onto the stack.

After iterating through all the temperatures, the result array will contain the
number of days until a warmer temperature is encountered for each day.
*/

// Time Complexity : O(n)
// Space Complexity: O(n)

// dailyTemperatures calculates for each day, the number of days until warmer temperature.
func dailyTemperatures(temperatures []int) []int {

	// Number of days until warmer temperature.
	waitDays := make([]int, len(temperatures)) // Preallocate for index management.

	// Implement a Decreasing Monotonic Stack of INDICES.
	var indexStack []int
	for today, currentTemp := range temperatures {

		// For every day in the past that is colder than today.
		for (len(indexStack) > 0) && (currentTemp > temperatures[indexStack[len(indexStack)-1]]) {

			// Calculate how many days ago was that.
			daysPast := indexStack[len(indexStack)-1] // index on top of stack.
			waitDays[daysPast] = today - daysPast

			// Discard this day as it's already been calculated above.
			indexStack = indexStack[:len(indexStack)-1]

		}
		// Push today on stack to calculate in the future.
		indexStack = append(indexStack, today)
	}
	return waitDays
}

///

func main() {

	var temperatures []int
	var want []int
	var got []int

	///

	fmt.Println("Test1:")
	temperatures = []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Printf("temperatures = %#v \n", temperatures)
	want = []int{1, 1, 4, 2, 1, 1, 0, 0}
	got = dailyTemperatures(temperatures)
	fmt.Printf("Want: %v \n Got: %v\n", want, got)
	fmt.Println()

}
