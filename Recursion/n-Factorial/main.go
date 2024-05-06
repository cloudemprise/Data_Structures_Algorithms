package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	tmp := n * factorial(n-1)
	fmt.Println(tmp)
	return tmp
}

func main() {

	val := 20
	fmt.Printf("%v factorial = %v", val, factorial(val))

}
