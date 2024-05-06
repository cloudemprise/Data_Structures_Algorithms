package main

import "fmt"

func main() {

	// Creating a 2D slice:
	rows, cols := 3, 4
	matrixA := make([][]int, rows)
	for i := range matrixA {
		matrixA[i] = make([]int, cols)
	} // matrixA[rows][cols]

	// Prints the value in full Go syntax:
	fmt.Println("A 2D Matrix in full Go syntax:")
	fmt.Printf("%#v \n", matrixA)

	fmt.Println()

	// Double-Loop Access Pattern:
	fmt.Println("A Double-Loop Access Pattern: ")
	counter := (rows * cols)
	for r := range matrixA {

		for c := range matrixA[r] {

			matrixA[r][c] = counter
			fmt.Print(matrixA[r][c], " ")

			counter--
		}
	}

	fmt.Println()
	fmt.Println()

	///

	// Creating a 2D array:
	const x, y = 3, 4
	matrixB := [x][y]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	// Single-Loop Access Pattern:
	fmt.Println("A Single-Loop Access Pattern: ")
	for i := 0; i < (x * y); i++ {
		fmt.Print(matrixB[i/y][i%y], " ")
	}

}
