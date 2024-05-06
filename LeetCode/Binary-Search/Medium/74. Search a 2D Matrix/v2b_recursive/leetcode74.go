package main

import "fmt"

/*
The time complexity is O(log(m*n)), where m is the number of rows in the matrix
 and n is the number of columns. This is because the algorithm uses binary
 search to find the target integer in the matrix, and each iteration of the
 binary search reduces the search space by half. Since the matrix has m*n
 elements, the number of iterations required to find the target integer is
 logarithmic in the total number of elements.

The space complexity of the solution is O(log(m x n)) because it is determined
 by the maximum number of recursive calls on the call stack.

Please note that the space complexity analysis assumes that the Go compiler does
 not perform any optimizations that would reduce the space usage of the
 recursive function due to it being a tail recursive function.
*/

// searchMatrix returns true if target exists within the given 2D matrix or false otherwise.
func searchMatrix(matrix [][]int, target int) bool {

	m, n := len(matrix), len(matrix[0]) // input matrix[m][n]

	// Recursive Closure:
	var search func(int, int) bool
	search = func(lo, hi int) bool {
		if lo > hi {
			return false
		}

		mid := (lo + hi) / 2
		switch {
		case target > matrix[mid/n][mid%n]:
			return search(mid+1, hi)
		case target < matrix[mid/n][mid%n]:
			return search(lo, mid-1)
		default:
			return true
		}
	}
	return search(0, (m*n)-1)
}

///

func main() {

	var matrix [][]int
	var target int

	var want bool
	var got bool

	///

	fmt.Println("TEST1:")
	matrix = [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	fmt.Printf("matrix = %v \n", matrix)
	target = 3
	fmt.Printf("target = %v \n", target)

	want = true
	got = searchMatrix(matrix, target)
	fmt.Printf("Want: %v \n Got: %v\n", want, got)
	fmt.Println()

	///

	fmt.Println("TEST2:")
	matrix = [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	fmt.Printf("matrix = %v\n", matrix)
	target = 13
	fmt.Printf("target = %v\n", target)

	want = false
	got = searchMatrix(matrix, target)
	fmt.Printf("Want: %v \n Got: %v\n", want, got)
	fmt.Println()

	///

	fmt.Println("TEST3:")
	matrix = [][]int{
		{1},
		{3},
	}
	fmt.Printf("matrix = %v\n", matrix)
	target = 2
	fmt.Printf("target = %v\n", target)

	want = false
	got = searchMatrix(matrix, target)
	fmt.Printf("Want: %v \n Got: %v\n", want, got)
	fmt.Println()

	///

}
