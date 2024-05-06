package main

import (
	"container/heap"
	"fmt"
	"math"
)

// > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > >

/*
2231. Largest Number After Digit Swaps by Parity
Easy

You are given a positive integer num. You may swap any two digits of num that
have the same parity (i.e. both odd digits or both even digits).

Return the largest possible value of num after any number of swaps.


Example 1:

Input: num = 1234
Output: 3412
Explanation: Swap the digit 3 with the digit 1, this results in the number 3214.
Swap the digit 2 with the digit 4, this results in the number 3412.
Note that there may be other sequences of swaps but it can be shown that 3412
is the largest possible number.
Also note that we may not swap the digit 4 with the digit 1 since they are of
different parities.

Example 2:

Input: num = 65875
Output: 87655
Explanation: Swap the digit 8 with the digit 6, this results in the number 85675.
Swap the first digit 5 with the digit 7, this results in the number 87655.
Note that there may be other sequences of swaps but it can be shown that 87655
is the largest possible number.


Constraints:

    1 <= num <= 109
*/

/*
Intuitive Understanding:
------------------------

*/

// < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < <

// Solution 3: More elegant using two Min-Heaps.

// Implement a Max-Heap:
type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(val any)      { (*h) = append((*h), val.(int)) }
func (h *minHeap) Pop() any {
	defer func() { (*h) = (*h)[:len(*h)-1] }()
	return (*h)[len(*h)-1]
}

// Separate digits and reconstitute using a Max-Heap
func largestInteger(num int) int {

	var numCopy int // a working copy

	// how many digits does the number have?
	numOfDigits := int(math.Log10(float64(num))) + 1

	//even and odd max-heaps
	var evens, odds minHeap

	// separate digits into evens & odds.
	numCopy = num
	for i := 0; i < numOfDigits; i++ {
		// read least significant digit
		digit := numCopy % 10
		if digit%2 == 0 {
			heap.Push(&evens, digit) // Push evens
		} else {
			heap.Push(&odds, digit) // Push odds
		}
		// remove least significant digit
		numCopy /= 10
	}

	// reconstitute the new number
	var numNew, value int
	numCopy = num
	for i := 0; i < numOfDigits; i++ {
		// read least significant digit
		digit := numCopy % 10
		if digit%2 == 0 {
			value = heap.Pop(&evens).(int) // Pop evens
		} else {
			value = heap.Pop(&odds).(int) // Pop odds
		}
		// remove least significant digit
		numCopy /= 10
		// include new most significant digit
		numNew += value * int(math.Pow10(i))
	}
	return numNew
}

//-------------------------------------

/* //-------------------------------------
// Solution 2: More elegant using two Max-Heaps.

// Implement a Max-Heap:
type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(val any)      { (*h) = append((*h), val.(int)) }
func (h *maxHeap) Pop() any {
	defer func() { (*h) = (*h)[:len(*h)-1] }()
	return (*h)[len(*h)-1]
}

// Separate digits and reconstitute using a Max-Heap
func largestInteger(num int) int {

	var numCopy int // a working copy

	// how many digits does the number have?
	numOfDigits := int(math.Log10(float64(num))) + 1

	//even and odd max-heaps
	var evens, odds maxHeap

	// separate digits into evens & odds.
	numCopy = num
	for i := 1; i <= numOfDigits; i++ {
		// generate a 10^x divisor.
		divisor := int(math.Pow10(numOfDigits - i))
		digit := numCopy / divisor // integer division
		if digit%2 == 0 {
			heap.Push(&evens, digit) // Push evens
		} else {
			heap.Push(&odds, digit) // Push odds
		}
		// remove most significant digit
		numCopy -= (digit * divisor)
	}

	// reconstitute the new number
	var numNew, integer int
	numCopy = num
	for i := 1; i <= numOfDigits; i++ {
		// generate a 10^x divisor.
		divisor := int(math.Pow10(numOfDigits - i))
		digit := numCopy / divisor // integer division
		if digit%2 == 0 {
			integer = heap.Pop(&evens).(int) // Pop evens
		} else {
			integer = heap.Pop(&odds).(int) // Pop odds
		}
		// remove old most significant digit
		numCopy -= (digit * divisor)
		// include new most significant digit
		numNew += (integer * divisor)
	}
	return numNew
} */

//-------------------------------------

/* //Solution 1: Quick & Dirty recovery from misinterpretation of the question.
// Made the mistake in thinking the question was refering to the indices of the
// digits not the parity of the digits themselves.

// Use math gymnastics + sort.Reverse
func largestInteger(num int) int {

	// how many digits does the number have?
	numOfDigits := int(math.Log10(float64(num))) + 1

	// separate digits into evens & odds and record their indices.
	var odds, evens []int
	indices := make([]int, numOfDigits)
	for i := 1; i <= numOfDigits; i++ {
		divisor := math.Pow10(numOfDigits - i)
		digit := int(num / int(divisor))
		if digit%2 == 0 {
			evens = append(evens, digit)
			indices[i-1] = 0 // indicate even index
		} else {
			odds = append(odds, digit)
			indices[i-1] = 1 // indicate odd index
		}
		num -= (digit * int(divisor))
	}
	// reverse sort each digit list
	sort.Sort(sort.Reverse(sort.IntSlice(odds)))
	sort.Sort(sort.Reverse(sort.IntSlice(evens)))

	// reconstitute the number
	var result int
	for i := 1; i <= numOfDigits; i++ {
		multiplier := int(math.Pow10(numOfDigits - i))
		if indices[i-1] == 0 {
			result = result + (multiplier * evens[0])
			evens = evens[1:] // reslice quick & dirty
		} else {
			result = result + (multiplier * odds[0])
			odds = odds[1:] // reslice quick & dirty
		}
	}
	return result
} */

//-------------------------------------

func main() {

	var num int

	num = 1234
	fmt.Println("Input number = ", num)
	fmt.Printf("Want: 3412 \t Got: %v\n", largestInteger(num))
	fmt.Println()

	num = 65875
	fmt.Println("Input number = ", num)
	fmt.Printf("Want: 87655 \t Got: %v\n", largestInteger(num))
	fmt.Println()

	num = 1
	fmt.Println("Input number = ", num)
	fmt.Printf("Want: 1 \t Got: %v\n", largestInteger(num))
	fmt.Println()

	num = 123 // answer = 427
	fmt.Println("Input number = ", num)
	fmt.Printf("Want: 321 \t Got: %v\n", largestInteger(num))
	fmt.Println()

	num = 247 // answer = 427
	fmt.Println("Input number = ", num)
	fmt.Printf("Want: 427 \t Got: %v\n", largestInteger(num))
	fmt.Println()

	num = 469 // answer = 427
	fmt.Println("Input number = ", num)
	fmt.Printf("Want: 649 \t Got: %v\n", largestInteger(num))
	fmt.Println()

}

//*-*
