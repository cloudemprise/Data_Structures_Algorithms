package main

import (
	"fmt"
	"math"
	"time"
)

/*
Time Complexity = O(n)
Space Complexity = O(n)
*/

// Using a hashmap
func majorityElement(nums []int) int {
	freq := make(map[int]int)
	var candidate int

	upperLimit := int(math.Ceil(float64(len(nums)) / 2))

	for _, num := range nums {
		freq[num]++
		if freq[num] > freq[candidate] {
			candidate = num
		}
		if freq[candidate] >= upperLimit { // only > 50% required.
			return candidate
		}
	}
	return int(^uint(0) >> 1) // No candidate.
}

///

func main() {

	var nums []int
	var result int
	var start time.Time
	var elapsed time.Duration

	nums = []int{3, 2, 3}
	fmt.Println("--- Input --- \n", nums)

	start = time.Now()
	result = majorityElement(nums)
	elapsed = time.Since(start)

	fmt.Printf("Want: 3 Got: %v\n", result)
	fmt.Printf("Execution Time : %v\n", elapsed)
	fmt.Println()

	///

	nums = []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println("--- Input --- \n", nums)

	start = time.Now()
	result = majorityElement(nums)
	elapsed = time.Since(start)

	fmt.Printf("Want: 2 Got: %v\n", result)
	fmt.Printf("Execution Time : %v\n", elapsed)
	fmt.Println()

	///

	nums = []int{1, 2, 1, 2, 0}
	fmt.Println("--- Input --- \n", nums)

	start = time.Now()
	result = majorityElement(nums)
	elapsed = time.Since(start)

	fmt.Printf("Want: %d Got: %v\n", int(^uint(0)>>1), result)
	fmt.Printf("Execution Time : %v\n", elapsed)
	fmt.Println()

}
