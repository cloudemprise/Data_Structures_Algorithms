package main

import "fmt"

// > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > >

/*
Comments
--------

Mininum number of moves to solve = 2^n-1

Typical Output:

Step  1: Move disk 1 from rod A to rod B
Step  2: Move disk 2 from rod A to rod C
Step  3: Move disk 1 from rod B to rod C
Step  4: Move disk 3 from rod A to rod B
Step  5: Move disk 1 from rod C to rod A
Step  6: Move disk 2 from rod C to rod B
Step  7: Move disk 1 from rod A to rod B
Step  8: Move disk 4 from rod A to rod C
Step  9: Move disk 1 from rod B to rod C
Step 10: Move disk 2 from rod B to rod A
Step 11: Move disk 1 from rod C to rod A
Step 12: Move disk 3 from rod B to rod C
Step 13: Move disk 1 from rod A to rod B
Step 14: Move disk 2 from rod A to rod C
Step 15: Move disk 1 from rod B to rod C

*/

// < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < <

func towerOfHanoi2(n, start, end int) {
	if n == 1 {
		fmt.Printf("Move disk %v from rod %d to rod %d\n", n, start, end)
	} else {
		aux := 6 - (start + end)
		towerOfHanoi2(n-1, start, aux)
		fmt.Printf("Move disk %v from rod %d to rod %d\n", n, start, end)
		towerOfHanoi2(n-1, aux, end)
	}
}

//-------------------------------------

type rod rune

var step int

// Time Complexity: O(2^n)
// Rod A = origin
// Rod B = auxillary
// Rod C = destination.
func towerOfHanoi(n int, A, C, B rod) {

	if n == 0 {
		return
	}

	towerOfHanoi(n-1, A, B, C)

	step++
	fmt.Printf("Step %2v: ", step)
	fmt.Printf("Move disk %v from rod %c to rod %c\n", n, A, C)

	towerOfHanoi(n-1, B, C, A)
}

//*-*

/* func towerOfHanoi(n int, A, C, B rod) {

	if n == 1 {
		step++
		fmt.Printf("Step %2v: ", step)
		fmt.Printf("Move disk %v from rod %c to rod %c\n", n, A, C)
	} else {
		towerOfHanoi(n-1, A, B, C)
		step++
		fmt.Printf("Step %2v: ", step)
		fmt.Printf("Move disk %v from rod %c to rod %c\n", n, A, C)
		towerOfHanoi(n-1, B, C, A)
	}
} */

func main() {

	towerOfHanoi(4, 'A', 'C', 'B')
	towerOfHanoi2(4, 1, 3)

}

//*-*
