package main

import (
	"container/heap"
)

// > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > >

/*
1464. Maximum Product of Two Elements in an Array
Easy

Given the array of integers nums, you will choose two different indices i and j
of that array. Return the maximum value of (nums[i]-1)*(nums[j]-1).


Example 1:

Input: nums = [3,4,5,2]
Output: 12
Explanation: If you choose the indices i=1 and j=2 (indexed from 0), you will
get the maximum value, that is:

(nums[1]-1)*(nums[2]-1) = (4-1)*(5-1) = 3*4 = 12.

Example 2:

Input: nums = [1,5,4,5]
Output: 16
Explanation: Choosing the indices i=1 and j=3 (indexed from 0), you will get
the maximum value of (5-1)*(5-1) = 16.

Example 3:

Input: nums = [3,7]
Output: 12


Constraints:

    2 <= nums.length <= 500
    1 <= nums[i] <= 10^3
*/

/*
Intuitive Understanding:
------------------------


*/

// < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < <

//-------------------------------------

// Sliding Window????
/* func maxProduct(nums []int) int {

	var first, second int

	for _, num := range nums {
		if num > first {
			second = first // keep copy of next max
			first = num
		} else if num > second {
			second = num
		}
	}
	return (first - 1) * (second - 1)

} */

//-------------------------------------

// Use stdlib container/heap to implement
// a Max-Heap ADT.

// maxHeap is a max-heap of ints.
type maxHeap []int

// implement sort.Interface
func (h maxHeap) Len() int {
	return len(h)
}

// implement sort.Interface
func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

// implement sort.Interface
func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// implement heap.Interface.
// reslice requires ptr-receiver.
func (h *maxHeap) Push(val any) {
	*h = append(*h, val.(int))
}

// implement heap.Interface.
// reslice requires ptr-receiver.
func (h *maxHeap) Pop() any {
	defer func() {
		*h = (*h)[:len(*h)-1]
	}()
	return (*h)[len(*h)-1]
}

func maxProduct(nums []int) int {
	h := maxHeap(nums) // type conversion.
	heap.Init(&h)
	return (heap.Pop(&h).(int) - 1) * (heap.Pop(&h).(int) - 1)
}

//-------------------------------------

func main() {

	PrintResult()
}

//*-*
