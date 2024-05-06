package main

import "fmt"

// transpose transposes an (n x m) matrix
func transpose(matrix [][]int) [][]int {

	var output [][]int
	for i := 0; i < len(matrix[0]); i++ {
		var tmp []int
		for j := 0; j < len(matrix); j++ {
			tmp = append(tmp, matrix[j][i])
		}
		output = append(output, tmp)
	}
	return output
}

func main() {

	input := [][]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}

	fmt.Println(transpose(input))
}
