package main

import (
	"fmt"
	"sort"
	"strconv"
)

// > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > >

/*
506. Relative Ranks
Easy
1.1K
53
Companies

You are given an integer array score of size n, where score[i] is the score of
the ith athlete in a competition. All the scores are guaranteed to be unique.

The athletes are placed based on their scores, where the 1st place athlete has
the highest score, the 2nd place athlete has the 2nd highest score, and so on.
The placement of each athlete determines their rank:

    The 1st place athlete's rank is "Gold Medal".
    The 2nd place athlete's rank is "Silver Medal".
    The 3rd place athlete's rank is "Bronze Medal".
    For the 4th place to the nth place athlete, their rank is their placement
	number (i.e., the xth place athlete's rank is "x").

Return an array answer of size n where answer[i] is the rank of the ith athlete.


Example 1:

Input: score = [5,4,3,2,1]
Output: ["Gold Medal","Silver Medal","Bronze Medal","4","5"]
Explanation: The placements are [1st, 2nd, 3rd, 4th, 5th].

Example 2:

Input: score = [10,3,8,9,4]
Output: ["Gold Medal","5","Bronze Medal","Silver Medal","4"]
Explanation: The placements are [1st, 5th, 3rd, 2nd, 4th].


Constraints:

    n == score.length
    1 <= n <= 104
    0 <= score[i] <= 106
    All the values in score are unique.

*/

/*
Intuitive Understanding:
------------------------

*/

// < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < <

/* // Solution 3: Simpler than using a Max-Heap.

// Using an athlete object comprising index & score while sorting rank
// using sort.Slice().
func findRelativeRanks(score []int) []string {

	type athlete struct {
		index int
		score int
	}
	// construct a slice of athlete objects.
	var athletes []athlete
	for i, v := range score {
		athletes = append(athletes, athlete{i, v})
	}
	// Implement the sort.Slice() function by stipulating the
	// Less() method for score maximum.
	sort.Slice(athletes, func(i, j int) bool {
		return athletes[i].score > athletes[j].score
	})
	// Construct output string from sorted slice.
	result := make([]string, len(score)) // need to preallocate some mem.
	for i, athlete := range athletes {
		switch i {
		case 0:
			result[athlete.index] = "Gold Medal"
		case 1:
			result[athlete.index] = "Silver Medal"
		case 2:
			result[athlete.index] = "Bronze Medal"
		default:
			result[athlete.index] = strconv.Itoa(i + 1)
		}
	}
	return result
} */

// -------------------------------------

// Solution 2: The simplest to understand.

// Using a map to record each athlete's index & score while sorting rank
// by way of sort.Sort(sort.Reverse(sort.IntSlice())).
// to determine placement order.
func findRelativeRanks(score []int) []string {

	athletes := make(map[int]int)
	for i, v := range score {
		athletes[v] = i
	}

	sort.Sort(sort.Reverse(sort.IntSlice(score)))
	result := make([]string, len(score)) // need to preallocate some mem.

	for i, v := range score {
		switch i {
		case 0:
			result[athletes[v]] = "Gold Medal"
		case 1:
			result[athletes[v]] = "Silver Medal"
		case 2:
			result[athletes[v]] = "Bronze Medal"
		default:
			result[athletes[v]] = strconv.Itoa(i + 1)
		}
	}
	return result
}

// -------------------------------------

/* // Solution 1:

// Implement a Max-Heap of each an athlete object comprising
// athlete's index & score.

type athlete struct {
	index int
	score int
}

type maxHeap []athlete

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i].score > h[j].score }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(val any)      { (*h) = append((*h), val.(athlete)) }
func (h *maxHeap) Pop() any {
	defer func() { (*h) = (*h)[:len(*h)-1] }()
	return (*h)[len(*h)-1]
}

// Push all athletes onto a heap & Pop while using index as
// placement into a string slice.
func findRelativeRanks(score []int) []string {

	var athletes maxHeap
	for i, v := range score {
		heap.Push(&athletes, athlete{i, v})
	}

	result := make([]string, len(athletes)) // need to preallocate some mem.

	for i := 1; athletes.Len() > 0; i++ {
		obj := heap.Pop(&athletes).(athlete)
		switch i {
		case 1:
			result[obj.index] = "Gold Medal"
		case 2:
			result[obj.index] = "Silver Medal"
		case 3:
			result[obj.index] = "Bronze Medal"
		default:
			result[obj.index] = strconv.Itoa(i)
		}
	}
	return result
} */

//-------------------------------------

func main() {

	/*
	   1. Value
	   2. Index
	   3. Relative Order
	*/

	score := []int{10, 3, 8, 9, 4}
	fmt.Printf("For input :\t\t %#v\n", score)
	fmt.Printf("Relative Ranks :\t %#v\n", findRelativeRanks(score))

}

//*-*
