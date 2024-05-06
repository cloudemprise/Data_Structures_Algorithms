package main

import "fmt"

// Allocate new slice for result
func runningSum(nums []int) []int {

	result := []int{nums[0]}
	for i, num := range nums {
		if i == 0 {
			continue
		}
		result = append(result, result[i-1]+num)
	}
	return result
}

///

func main() {

	input := []int{1, 2, 3, 4}
	fmt.Println(runningSum(input))

}
