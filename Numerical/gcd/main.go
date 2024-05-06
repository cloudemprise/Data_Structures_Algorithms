package main

import "fmt"

// > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > >

/*
Greatest Common Divisor
*/

// < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < <

/* // Euclid's Algo.
func gcd(m, n int) int {
	if m == 0 {
		return n
	} else if m > n {
		return gcd(n, m)
	} else {
		return gcd(m, n-m)
	}
} */

// More efficient version.
func gcd(m, n int) int {
	if m == 0 {
		return n
	} else {
		return gcd(n%m, m)
	}
}

//-------------------------------------

func main() {

	m := 24
	n := 20
	fmt.Printf("gcd(%v,%v) = %v\n", m, n, gcd(m, n))

}

//*-*
