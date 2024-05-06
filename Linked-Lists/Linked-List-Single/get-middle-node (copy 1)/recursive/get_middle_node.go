package main

import "fmt"

//

/*
The first and second middle nodes of a singly linked list refer to the nodes
that are located at the middle position(s) of the list.

When determining the middle node(s) of a singly linked list, it involves
identifying the node(s) that divide the list into two equal parts or close to
equal parts based on whether the total number of nodes is odd or even.

- If the linked list has an odd number of nodes, the ONLY middle node is the
	one that is exactly in the middle.

- But if the linked list has an even number of nodes then there exists two
	middle nodes. The second middle node is considered as the one closer to the
	end of the list among the two middle nodes.

Depending on your application you might choose one or the other.
*/

// get1stMiddleNode returns the 1st middle node of a singly linked list.
func get1stMiddleNode(head *Node) *Node {
	return find1stMiddle(head, head)
}

// find1stMiddle is a helper function that recursively finds the first middle
// node of a singly linked list.
/* func find1stMiddle(slow, fast *Node) *Node {
	if fast == nil || fast.Next == nil {
		return slow
	}
	if fast.Next.Next == nil {
		return slow
	}
	return find1stMiddle(slow.Next, fast.Next.Next)
} */
func find1stMiddle(slow, fast *Node) *Node {
	if fast == nil || fast.Next == nil || fast.Next.Next == nil {
		return slow
	}
	/* if fast.Next.Next == nil {
		return slow
	} */
	return find1stMiddle(slow.Next, fast.Next.Next)
}

///

// get2ndMiddleNode returns the 2nd middle node of a singly linked list.
func get2ndMiddleNode(head *Node) *Node {
	return find2ndMiddle(head, head)
}

// find2ndMiddle is a helper function that recursively finds the second middle
// node of a singly linked list.
func find2ndMiddle(slow, fast *Node) *Node {
	if fast == nil || fast.Next == nil {
		return slow
	}
	return find2ndMiddle(slow.Next, fast.Next.Next)
}

///

func main() {

	fmt.Println("Given a linked list with five (odd) nodes:")
	fmt.Println("------------------------------------------")

	oddList := &Node{
		Val: 1,
		Next: &Node{
			Val: 2,
			Next: &Node{
				Val: 3,
				Next: &Node{
					Val: 4,
					Next: &Node{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	fmt.Println("get1stMiddleNode() =", get1stMiddleNode(oddList).Val)
	fmt.Println("get2ndMiddleNode() =", get2ndMiddleNode(oddList).Val)
	fmt.Println("find1stMiddle() =", find1stMiddle(oddList, oddList).Val)
	fmt.Println("find2ndMiddle() =", find2ndMiddle(oddList, oddList).Val)
	fmt.Println()

	///

	fmt.Println("Given a linked list with six (even) nodes:")
	fmt.Println("------------------------------------------")

	evenList := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node5 := &Node{Val: 5}
	node6 := &Node{Val: 6}

	evenList.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6

	fmt.Println("get1stMiddleNode() =", get1stMiddleNode(evenList).Val)
	fmt.Println("get2ndMiddleNode() =", get2ndMiddleNode(evenList).Val)
	fmt.Println("find1stMiddle() =", find1stMiddle(evenList, evenList).Val)
	fmt.Println("find2ndMiddle() =", find2ndMiddle(evenList, evenList).Val)
	fmt.Println()
}
