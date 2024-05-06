package main

import (
	"fmt"
	"strconv"
)

/*
Solution 2 description.
-----------------------

We're looking to build three grids using the number in the Sudoku(digit) as each
position in the following three matrices. Note the matrix access patterns are
not the same as in Solution 1, i.e.

for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			digit := sudokuExtract
			horizontal[row][digit] = indicator
			vertical[digit][col] = indicator
			gridIndex = (col/3) + 3*(row/3)
			inner[gridIndex][digit] = indicator
	}
}

1. Horizontal: Sudoku digit = digits[1...9]
 --- --- --- --- --- --- --- --- ---
| 1   2   3   4   5   6   7   8   9 |
 --- --- --- --- --- --- --- --- ---
| 1   2   3   4   5   6   7   8   9 |
 --- --- --- --- --- --- --- --- ---
| 1   2   3   4   5   6   7   8   9 |
 --- --- --- --- --- --- --- --- ---
| 1   2   3   4   5   6   7   8   9 |
 --- --- --- --- --- --- --- --- ---
| 1   2   3   4   5   6   7   8   9 |
 --- --- --- --- --- --- --- --- ---
| 1   2   3   4   5   6   7   8   9 |
 --- --- --- --- --- --- --- --- ---
| 1   2   3   4   5   6   7   8   9 |
 --- --- --- --- --- --- --- --- ---
| 1   2   3   4   5   6   7   8   9 |
 --- --- --- --- --- --- --- --- ---
| 1   2   3   4   5   6   7   8   9 |
 --- --- --- --- --- --- --- --- ---


2. Vertical: Sudoku digit = [1...9]
 --- --- --- --- --- --- --- --- ---
| 1 | 1 | 1 | 1 | 1 | 1 | 1 | 1 | 1 |
| 2 | 2 | 2 | 2 | 2 | 2 | 2 | 2 | 2 |
| 3 | 3 | 3 | 3 | 3 | 3 | 3 | 3 | 3 |
| 4 | 4 | 4 | 4 | 4 | 4 | 4 | 4 | 4 |
| 5 | 5 | 5 | 5 | 5 | 5 | 5 | 5 | 5 |
| 6 | 6 | 6 | 6 | 6 | 6 | 6 | 6 | 6 |
| 7 | 7 | 7 | 7 | 7 | 7 | 7 | 7 | 7 |
| 8 | 8 | 8 | 8 | 8 | 8 | 8 | 8 | 8 |
| 9 | 9 | 9 | 9 | 9 | 9 | 9 | 9 | 9 |
 --- --- --- --- --- --- --- --- ---

3. Inner: Sudoku digit = [1...9]
 --- --- --- --- --- --- --- --- ---
| 1   2   3 | 1   2   3 | 1   2   3 |
| 4   5   6 | 4   5   6 | 4   5   6 |
| 7   8   9 | 7   8   9 | 7   8   9 |
 --- --- --- --- --- --- --- --- ---
| 1   2   3 | 1   2   3 | 1   2   3 |
| 4   5   6 | 4   5   6 | 4   5   6 |
| 7   8   9 | 7   8   9 | 7   8   9 |
 --- --- --- --- --- --- --- --- ---
| 1   2   3 | 1   2   3 | 1   2   3 |
| 4   5   6 | 4   5   6 | 4   5   6 |
| 7   8   9 | 7   8   9 | 7   8   9 |
 --- --- --- --- --- --- --- --- ---
*/

// Solution v2a:
func isValidSudoku(board [][]byte) bool {

	// Three 9x9 grid permutations of digits[1...9] to check.
	// Arrays are preinitialised, unlike slices.
	var horizontal, vertical, inner [9][9]bool

	// Visit each cell only once.
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {

			// Read a site.
			sudokuVal := board[row][col]

			// Only process valid digits, ignore all else.
			if sudokuVal >= '1' && sudokuVal <= '9' {

				digit := int(sudokuVal - '0')
				digit-- // Decrement to get index[0...8] as per arrays.

				// Check if current sites have already been visited.
				gridIndex := (col / 3) + 3*(row/3) // Calculate inner grid position.
				if horizontal[row][digit] || vertical[digit][col] || inner[gridIndex][digit] {
					return false // Duplicate found.
				}

				// Otherwise mark sites as visited.
				horizontal[row][digit] = true
				vertical[digit][col] = true
				inner[gridIndex][digit] = true
			}
		}
	}
	return true // all checks made.
}

///

func main() {

	//printAccessPatterns()

	var input [][]byte

	input = [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	// visualize input via access pattern 1, horizontal:
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			digit, _ := strconv.Atoi(string(input[row][col]))
			fmt.Printf("%v ", digit)
		}
		fmt.Println()
	}
	fmt.Printf("Want: true \t Got: %v\n", isValidSudoku(input))
	fmt.Println()

	//*-*

	input = [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	// visualize input via access pattern 1, horizontal:
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			digit, _ := strconv.Atoi(string(input[row][col]))
			fmt.Printf("%v ", digit)
		}
		fmt.Println()
	}
	fmt.Printf("Want: false \t Got: %v\n", isValidSudoku(input))
	fmt.Println()

	//*-*

	input = [][]byte{
		{'.', '.', '4', '.', '.', '.', '6', '3', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'5', '.', '.', '.', '.', '.', '.', '9', '.'},
		{'.', '.', '.', '5', '6', '.', '.', '.', '.'},
		{'4', '.', '3', '.', '.', '.', '.', '.', '1'},
		{'.', '.', '.', '7', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '5', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	}

	// visualize input via access pattern 1, horizontal:
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			digit, _ := strconv.Atoi(string(input[row][col]))
			fmt.Printf("%v ", digit)
		}
		fmt.Println()
	}
	fmt.Printf("Want: false \t Got: %v\n", isValidSudoku(input))
	fmt.Println()

}

//*-*

func printAccessPatterns() {

	//
	const (
		lenRow = 9
		lenCol = 9
	)

	var board [lenRow][lenCol]int
	var counter int

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			counter++
			board[row][col] = row*9 + (col + 1)
		}
	}

	// access pattern 1, horizontal:
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			fmt.Printf("%02v ", board[row][col])
		}
		fmt.Println()
	}
	fmt.Println()

	// access pattern 2, vertical:
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			fmt.Printf("%02v ", board[col][row])
		}
		fmt.Println()
	}
	fmt.Println()

	// access pattern 3, inner square.
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			column := ((col % 3) + (3 * row)) % 9
			line := (col / 3) + 3*(row/3)
			fmt.Printf("%02v ", board[line][column])
		}
		fmt.Println()
	}
	fmt.Println()
}

//*-*
