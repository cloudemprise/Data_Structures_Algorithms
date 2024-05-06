package main

import "fmt"

/*
The time complexity is O(log(m*n)), where m is the number of rows in the matrix
 and n is the number of columns. This is because the algorithm uses binary
 search to find the target integer in the matrix, and each iteration of the
 binary search reduces the search space by half. Since the matrix has m*n
 elements, the number of iterations required to find the target integer is
 logarithmic in the total number of elements.

The space complexity of the solution is O(1) because it uses a constant amount
 of extra space to store the indices and variables needed for the binary search.
*/

// searchMatrix returns true if target exists within the given 2D matrix or false otherwise.
func searchMatrix(matrix [][]int, target int) bool {

	m, n := len(matrix), len(matrix[0]) // input matrix[m][n]

	lo, hi := 0, (m*n)-1 // index manipulation
	for lo <= hi {

		mid := (lo + hi) / 2
		switch {
		case target > matrix[mid/n][mid%n]:
			lo = mid + 1
		case target < matrix[mid/n][mid%n]:
			hi = mid - 1
		case target == matrix[mid/n][mid%n]:
			return true
		}
	}
	return false
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
