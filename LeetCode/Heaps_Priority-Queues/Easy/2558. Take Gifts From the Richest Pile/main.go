package main

import (
	"container/heap"
	"fmt"
	"math"
)

// > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > >

/*
2558. Take Gifts From the Richest Pile
Easy

You are given an integer array gifts denoting the number of gifts in various
piles. Every second, you do the following:

    Choose the pile with the maximum number of gifts.
    If there is more than one pile with the maximum number of gifts, choose any.
    Leave behind the floor of the square root of the number of gifts in the
	pile. Take the rest of the gifts.

Return the number of gifts remaining after k seconds.


Example 1:

Input: gifts = [25,64,9,4,100], k = 4
Output: 29
Explanation:
The gifts are taken in the following way:
- In the first second, the last pile is chosen and 10 gifts are left behind.
- Then the second pile is chosen and 8 gifts are left behind.
- After that the first pile is chosen and 5 gifts are left behind.
- Finally, the last pile is chosen again and 3 gifts are left behind.
The final remaining gifts are [5,8,9,4,3], so the total number of gifts
remaining is 29.

Example 2:

Input: gifts = [1,1,1,1], k = 4
Output: 4
Explanation:
In this case, regardless which pile you choose, you have to leave behind 1 gift
in each pile.
That is, you can't take any pile with you.
So, the total gifts remaining are 4.


Constraints:

    1 <= gifts.length <= 103
    1 <= gifts[i] <= 109
    1 <= k <= 103
*/

/*
Intuitive Understanding:
------------------------

*/

// < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < <

//-------------------------------------

/* // Solution 2: Using nested loops (inefficient)

func pickGifts(gifts []int, k int) int64 {

	for k := k; k > 0; k-- {
		var max, idx int
		for i, v := range gifts {
			if v > max {
				max = v
				idx = i
			}
		}
		gifts[idx] = int(math.Sqrt(float64(max)))
	}
	var result int
	for _, v := range gifts {
		result += v
	}
	return int64(result)
} */

// -------------------------------------

// Solution 1: Using a Max-Heap + simple logic.

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

// Push + heapify + Pop*logic + add
func pickGifts(gifts []int, k int) int64 {

	h := maxHeap(gifts) // type conversion
	heap.Init(&h)       // heapify
	// loop k times & implement logic
	for i := k; i > 0; i-- {
		max := heap.Pop(&h).(int) // type assertion
		leftover := int(math.Sqrt(float64(max)))
		if leftover > 0 {
			heap.Push(&h, leftover)
		}
	}
	// calculate result
	var result int
	for _, v := range h {
		result += v
	}
	return int64(result)
}

//-------------------------------------

func main() {

	gifts := []int{25, 64, 9, 4, 100}
	k := 4
	fmt.Println("Output = ", pickGifts(gifts, k))

}

//*-*
