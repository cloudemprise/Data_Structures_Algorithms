package main

import "fmt"

/*
Explanation:

- We use XOR to find the duplicate number.
- We initialize a variable result to 0 and XOR all the elements in the array with
	it.
- We then XOR all the numbers from 1 to n-1 with result, where n is the length of
	the array.
- The resulting value of result will be the duplicate number.
- This is because XORing two identical numbers results in 0, so all the
	non-duplicate numbers cancel out, leaving only the duplicate number.

	This solution has a time complexity of O(n) and a space complexity of O(1).
*/

/*
This doesn't work if there is more than one repeat of the same number, eg:
input = []int{1,2,3,4,2,2}
*/

func findDuplicate(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result ^= nums[i]
	}
	for i := 1; i < len(nums); i++ {
		result ^= i
	}
	return result
}

///

func main() {

	var input []int

	fmt.Println("Test1:")
	input = []int{1, 3, 4, 2, 2}
	fmt.Printf("input = %v\n", input)
	fmt.Printf("%v\n", findDuplicate(input))
	fmt.Println()

	//*-*

	fmt.Println("Test2:")
	input = []int{3, 1, 3, 4, 2}
	fmt.Printf("input = %v\n", input)
	fmt.Printf("%v\n", findDuplicate(input))
	fmt.Println()

	//*-*

	fmt.Println("Test3:")
	input = []int{8, 7, 1, 10, 17, 15, 18, 11, 16, 9, 19, 12, 5, 14, 3, 4, 2, 13, 6, 18}
	fmt.Printf("input = %v\n", input)
	fmt.Printf("%v\n", findDuplicate(input))
	fmt.Println()

	//*-*

}

//*-*
