package main

import (
	"container/heap"
	"fmt"
)

// > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > >

/*
703. Kth Largest Element in a Stream
Easy

Design a class to find the kth largest element in a stream. Note that it is the
kth largest element in the sorted order, not the kth distinct element.

Implement KthLargest class:

    KthLargest(int k, int[] nums) Initializes the object with the integer k and
	the stream of integers nums.
    int add(int val) Appends the integer val to the stream and returns the
	element representing the kth largest element in the stream.


Example 1:

Input
["KthLargest", "add", "add", "add", "add", "add"]
[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
Output
[null, 4, 5, 5, 8, 8]

Explanation
KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
kthLargest.add(3);   // return 4
kthLargest.add(5);   // return 5
kthLargest.add(10);  // return 5
kthLargest.add(9);   // return 8
kthLargest.add(4);   // return 8


Constraints:

    1 <= k <= 104
    0 <= nums.length <= 104
    -104 <= nums[i] <= 104
    -104 <= val <= 104
    At most 104 calls will be made to add.
    It is guaranteed that there will be at least k elements in the array when
	you search for the kth element.
*/

/*
Intuitive Understanding:
------------------------


*/

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

// < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < <

//-------------------------------------

/* type KthLargest struct {
	k    int
	nums []int
}

//*-*

func Constructor(k int, nums []int) KthLargest {
	return KthLargest{
		k:    k,
		nums: nums,
	}
}

//*-*

func (this *KthLargest) Add(val int) int {
	this.nums = append(this.nums, val)
	// descending order
	sort.Sort(sort.Reverse(sort.IntSlice(this.nums)))
	return this.nums[this.k-1]
} */

//-------------------------------------

// Implement a Min-Heap:
type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(val any)      { (*h) = append((*h), val.(int)) }
func (h *minHeap) Pop() any {
	defer func() { (*h) = (*h)[:len(*h)-1] }()
	return (*h)[len(*h)-1]
}

//*-*

type KthLargest struct {
	k    int
	nums minHeap
}

//*-*

func Constructor(k int, nums []int) KthLargest {
	h := minHeap(nums) // type conversion
	heap.Init(&h)
	for h.Len() > k { // don't need > k values
		heap.Pop(&h)
	}
	return KthLargest{k, h}
}

//*-*

func (this *KthLargest) Add(val int) int {
	if this.nums.Len() < this.k {
		heap.Push(&this.nums, val)
	} else if val > this.nums[0] {
		heap.Push(&this.nums, val)
		heap.Pop(&this.nums)
	}
	return this.nums[0]
}

//-------------------------------------

func main() {

	obj := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println("Initial Object = ", obj)

	fmt.Println()
	fmt.Println("Add 3 : Want 4 : Got ", obj.Add(3))
	fmt.Println(obj)

	fmt.Println()
	fmt.Println("Add 5 : Want 5 : Got ", obj.Add(5))
	fmt.Println(obj)

	fmt.Println()
	fmt.Println("Add 10 : Want 5 : Got ", obj.Add(10))
	fmt.Println(obj)

	fmt.Println()
	fmt.Println("Add 9 : Want 8 : Got ", obj.Add(10))
	fmt.Println(obj)

	fmt.Println()
	fmt.Println("Add 4 : Want 8 : Got ", obj.Add(4))
	fmt.Println(obj)

}

//*-*
