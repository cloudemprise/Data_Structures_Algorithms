package main

import "testing"

func TestHasCycle(t *testing.T) {
	// Create a linked list with a cycle
	head1 := &Node{Val: 3}
	node1 := &Node{Val: 2}
	node2 := &Node{Val: 0}
	node3 := &Node{Val: -4}
	head1.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node1

	// Create a linked list without a cycle
	head2 := &Node{Val: 1}
	node4 := &Node{Val: 2}
	node5 := &Node{Val: 3}
	node6 := &Node{Val: 4}
	head2.Next = node4
	node4.Next = node5
	node5.Next = node6

	testCases := []struct {
		name   string
		input  *Node
		output bool
	}{
		{
			name:   "cycle exists",
			input:  head1,
			output: true,
		},
		{
			name:   "cycle does not exist",
			input:  head2,
			output: false,
		},
		{
			name:   "single node",
			input:  &Node{Val: 1},
			output: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := hasCycle(tc.input)
			if result != tc.output {
				t.Errorf("hasCycle(%v) = %v; want %v", tc.input, result, tc.output)
			}
		})
	}
}
