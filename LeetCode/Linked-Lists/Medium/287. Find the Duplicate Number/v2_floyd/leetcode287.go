package main

import "fmt"

/*
Explanation:

- We use the Floyd's Tortoise and Hare (Cycle Detection) algorithm to detect the
	cycle in the given array.
- We start by initializing two pointers, slow and fast, to the first element of
	the array.
- We then move the slow pointer one step at a time and the fast pointer two
	steps at a time until they meet AT SOME POINT INSIDE THE CYCLE.
- Once we have detected the cycle, we initialize two more pointers, ptr1 and
	ptr2, to the first element of the array and the point where the slow and
	fast pointers met, respectively.
- We then move both pointers one step at a time until they meet at the entry
	point of the cycle, which is the duplicate number we are looking for.
- Finally, we return the duplicate number.

This solution has a time complexity of O(n) and a space complexity of O(1).
*/

/*
Reference:

https://en.wikipedia.org/wiki/Cycle_detection#Tortoise_and_hare
https://leetcode.com/problems/find-the-duplicate-number/editorial/
https://www.youtube.com/watch?v=wjYnzkAhcNk
*/

func findDuplicate(nums []int) int {

	// Phase 1: Find the intersection point.
	slow, fast := nums[0], nums[0]
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	// Meet at the Cyclic Entrance.
	ptr1, ptr2 := nums[0], slow
	for ptr1 != ptr2 {
		ptr1 = nums[ptr1]
		ptr2 = nums[ptr2]
	}
	return ptr1
}

///

func main() {

	var input []int

	/* fmt.Println("Test0:")
	input = []int{1, 9, 1, 8, 7, 6, 5, 4, 3, 2}
	fmt.Printf("input = %v\n", input)
	fmt.Printf("%v\n", findDuplicate(input))
	fmt.Println() */

	//*-*

	fmt.Println("Test1:")
	input = []int{1, 3, 4, 2, 2}
	fmt.Printf("input = %v\n", input)
	fmt.Printf("Want: 2 \t Got: %v\n", findDuplicate(input))
	fmt.Println()

	//*-*

	fmt.Println("Test2:")
	input = []int{3, 1, 3, 4, 2}
	fmt.Printf("input = %v\n", input)
	fmt.Printf("Want: 3 \t Got: %v\n", findDuplicate(input))
	fmt.Println()

	//*-*

	fmt.Println("Test3:")
	input = []int{8, 7, 1, 10, 17, 15, 18, 11, 16, 9, 19, 12, 5, 14, 3, 4, 2, 13, 6, 18}
	fmt.Printf("input = %v\n", input)
	fmt.Printf("Want: 18 \t Got: %v\n", findDuplicate(input))
	fmt.Println()

	//*-*

}

//*-*
