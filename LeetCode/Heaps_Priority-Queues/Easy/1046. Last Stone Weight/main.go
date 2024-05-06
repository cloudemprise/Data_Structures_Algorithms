package main

import (
	"container/heap"
	"fmt"
)

// > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > >

/*
1046. Last Stone Weight
Easy

You are given an array of integers stones where stones[i] is the weight of the
ith stone.

We are playing a game with the stones. On each turn, we choose the heaviest two
stones and smash them together. Suppose the heaviest two stones have weights x
and y with x <= y. The result of this smash is:

    If x == y, both stones are destroyed, and
    If x != y, the stone of weight x is destroyed, and the stone of weight y
	has new weight y - x.

At the end of the game, there is at most one stone left.

Return the weight of the last remaining stone. If there are no stones left,
return 0.


Example 1:

Input: stones = [2,7,4,1,8,1]
Output: 1
Explanation:
We combine 7 and 8 to get 1 so the array converts to [2,4,1,1,1] then,
we combine 2 and 4 to get 2 so the array converts to [2,1,1,1] then,
we combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
we combine 1 and 1 to get 0 so the array converts to [1] then that's the value
of the last stone.

Example 2:

Input: stones = [1]
Output: 1


Constraints:

    1 <= stones.length <= 30
    1 <= stones[i] <= 1000

*/

/*
Intuitive Understanding:
------------------------

*/

// < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < <

//-------------------------------------

// -------------------------------------

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

func lastStoneWeight(stones []int) int {

	game := maxHeap(stones) // type conversion.
	heap.Init(&game)        // heapify.

	for game.Len() > 1 {
		y := heap.Pop(&game).(int) // type cast.
		x := heap.Pop(&game).(int) // type cast.
		if smash := y - x; smash > 0 {
			heap.Push(&game, smash)
		}
	}

	if game.Len() > 0 {
		return game[0]
	} else {
		return 0
	}

}

//-------------------------------------

func main() {

	stones := []int{2, 7, 4, 1, 8, 1}
	fmt.Println("For input : ", stones, "Game ending :", lastStoneWeight(stones))

}

//*-*
