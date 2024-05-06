package main

import (
	"fmt"
	"strconv"
)

/*
Solution 1 description.
-----------------------
We're looking to analyse all 3x9=27 Sudoku solutions, (horizontal, vertical,
inner grid) iteratively whereby we search for duplicate values within these
solution groupings. Note the matrix access patterns are not the same as in
Solution 2, i.e.

for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			val := sudokuExtract
			horizontal[row][col] = indicator
			vertical[col][row] = indicator
			gridRow = (col/3) + 3*(row/3)
			gridCol = (col%3) + 3*(row%3)
			inner[gridRow][gridCol] = indicator
	}
}

1. Horizontal: Sudoku val = digits[1...9]
 --- --- --- --- --- --- --- --- ---
|  1   2   3   4   5   6   7   8   9 |
 --- --- --- --- --- --- --- --- ---
| 10  11  12  13  14  15  16  17  18 |
 --- --- --- --- --- --- --- --- ---
| 19  20  21  22  23  24  25  26  27 |
 --- --- --- --- --- --- --- --- ---
| 28  29  30  31  32  33  34  35  36 |
 --- --- --- --- --- --- --- --- ---
| 37  38  39  40  41  42  43  44  45 |
 --- --- --- --- --- --- --- --- ---
| 46  47  48  49  50  51  52  53  54 |
 --- --- --- --- --- --- --- --- ---
| 55  56  57  58  59  60  61  62  63 |
 --- --- --- --- --- --- --- --- ---
| 64  65  66  67  68  69  70  71  72 |
 --- --- --- --- --- --- --- --- ---
| 73  74  75  76  77  78  79  80  81 |
 --- --- --- --- --- --- --- --- ---


2. Vertical: Sudoku val = digits[1...9]
 --- --- --- --- --- --- --- --- ---
|  1 | 10 | 19 | 28 | 37 | 46 | 55 | 64 | 73 |
|  2 | 11 | 20 | 29 | 38 | 47 | 56 | 65 | 74 |
|  3 | 12 | 21 | 30 | 39 | 48 | 57 | 66 | 75 |
|  4 | 13 | 22 | 31 | 40 | 49 | 58 | 67 | 76 |
|  5 | 14 | 23 | 32 | 41 | 50 | 59 | 68 | 77 |
|  6 | 15 | 24 | 33 | 42 | 51 | 60 | 69 | 78 |
|  7 | 16 | 25 | 34 | 43 | 52 | 61 | 70 | 79 |
|  8 | 17 | 26 | 35 | 44 | 53 | 62 | 71 | 80 |
|  9 | 18 | 27 | 36 | 45 | 54 | 63 | 72 | 81 |
 --- --- --- --- --- --- --- --- ---

3. Inner: Sudoku val = digits[1...9]
 --- --- --- --- --- --- --- --- ---
|  1   2   3 | 10  11  12 | 19  20  21 |
|  4   5   6 | 13  14  15 | 22  23  24 |
|  7   8   9 | 16  17  18 | 25  26  27 |
 --- --- --- --- --- --- --- --- ---
| 28  29  30 | 37  38  39 | 46  47  48 |
| 31  32  33 | 40  41  42 | 49  50  51 |
| 34  35  36 | 43  44  45 | 52  53  54 |
 --- --- --- --- --- --- --- --- ---
| 55  56  57 | 64  65  66 | 73  74  75 |
| 58  59  60 | 67  68  69 | 76  77  78 |
| 61  62  63 | 70  71  72 | 79  80  81 |
 --- --- --- --- --- --- --- --- ---

*/

// Solution v1b: (Most Efficient)
func isValidSudoku(board [][]byte) bool {

	// Visit each cell only once.
	for row := 0; row < 9; row++ {
		// Reuse this to peek into Sudoku matrix.
		var digit byte

		// Each iteration we build a map for each of the 3
		// solution types using digits[1...9]
		// Remember maps need preallocation.
		horizontal := make(map[byte]struct{})
		vertical := make(map[byte]struct{})
		inner := make(map[byte]struct{})

		// Each loop performs 3x9 checks.
		for col := 0; col < 9; col++ {

			// Check 1: horizontal
			digit = board[row][col] // Read particular site.
			// Only process valid ASCII digits, ignore all else.
			if digit >= '1' && digit <= '9' {
				// Check if value is a duplicate.
				if _, ok := horizontal[digit]; ok {
					return false
				}
				// Otherwise mark it as seen.
				horizontal[digit] = struct{}{}
			}

			// Check 2: vertical.
			digit = board[col][row] // Read particular site.
			// Only process valid ASCII digits, ignore all else.
			if digit >= '1' && digit <= '9' {
				// Check if value is a duplicate.
				if _, ok := vertical[digit]; ok {
					return false
				}
				// Otherwise mark it as seen.
				vertical[digit] = struct{}{}
			}

			// Check 3: inner grid + calculate respective indices.
			r := (col / 3) + 3*(row/3)
			//c := ((col % 3) + (3 * row)) % 9
			c := (col % 3) + 3*(row%3) // Simplified easier to remember.
			digit = board[r][c]        // Read particular site.
			// Only process valid ASCII digits, ignore all else.
			if digit >= '1' && digit <= '9' {
				// Check if value is a duplicate.
				if _, ok := inner[digit]; ok {
					return false
				}
				// Otherwise mark it as seen.
				inner[digit] = struct{}{}
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
			val, _ := strconv.Atoi(string(input[row][col]))
			fmt.Printf("%v ", val)
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
			val, _ := strconv.Atoi(string(input[row][col]))
			fmt.Printf("%v ", val)
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
			val, _ := strconv.Atoi(string(input[row][col]))
			fmt.Printf("%v ", val)
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
