package main

import (
	"fmt"
)

func PrintResult() {

	/* var nums []int
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 7; i++ {
		nums = append(nums, r.Intn(9))
	} */

	nums := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Printf("For input slice : ")

	for _, num := range nums {
		fmt.Printf(" %v ", num)
	}
	fmt.Println()
	fmt.Printf("Answer is : %v", maxProduct(nums))

	//*-*
	fmt.Printf("\n//-------------------------------------\n\n")
	//*-*

}
