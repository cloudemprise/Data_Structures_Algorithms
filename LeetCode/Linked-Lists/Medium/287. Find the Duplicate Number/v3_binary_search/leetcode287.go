package main

import "fmt"

/*
Explanation:

- We use binary search to find the duplicate number.
- We initialize the lo and hi pointers to 1 and n-1, respectively, where n
	is the length of the array.
- We then calculate the mid point and count the number of elements in the array
	that are less than or equal to the mid point.
- If the count is less than or equal to the mid point, then the duplicate number
	must be in the hi half of the array, so we update the lo pointer to
	mid+1.
- Otherwise, the duplicate number must be in the lo half of the array, so we
	update the hi pointer to mid.
- We repeat this process until the lo and hi pointers converge to the
	duplicate number, which is returned.

This solution has a time complexity of O(nlogn) and a space complexity of O(1).
*/

func findDuplicate(nums []int) int {
	lo, hi := 1, len(nums)-1

	for lo < hi {
		mid := (hi + lo) / 2

		count := 0
		for _, num := range nums {
			if num <= mid {
				count++
			}
		}

		if count <= mid {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
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
	input = []int{8, 7, 1, 10, 17, 15, 18, 11, 16, 9, 19, 12, 5, 14, 3, 4, 2, 13, 18, 18}
	fmt.Printf("input = %v\n", input)
	fmt.Printf("%v\n", findDuplicate(input))
	fmt.Println()

	//*-*

}

//*-*
