package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Tree struct {
	root *TreeNode
}

// inorderTraversal performs a recursive in-order traversal of a binary-tree
// using a helper function that (pointer semantics) passes a `results` slice
// through the recursive stack.
func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	inorder(root, &result)
	return result
}

func inorder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, result)
	*result = append(*result, root.Val)
	inorder(root.Right, result)
}

///

func main() {

	aTree := &Tree{}
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	for _, num := range nums {
		aTree.insertBreadthFirst(num)
	}

	fmt.Println("Recursive In-Order Traversal:")
	fmt.Println(inorderTraversal(aTree.root))
}

//*-*

// insertBreadthFirst iteratively construct a binary-tree for test.
func (t *Tree) insertBreadthFirst(d int) {
	if t.root == nil {
		t.root = &TreeNode{Val: d}
	} else {

		var queue []*TreeNode
		queue = append(queue, t.root)
		for len(queue) > 0 {
			// DeQueue
			current := queue[0]
			queue = queue[1:]

			if current.Left == nil { // move left
				current.Left = &TreeNode{Val: d}
				break
			} else {
				queue = append(queue, current.Left)
			}

			if current.Right == nil { // move right
				current.Right = &TreeNode{Val: d}
				break
			} else {
				queue = append(queue, current.Right)
			}
		}
	}
}

/* // insertBinarySearch recursive helper function to construct a binary-search-tree for test.
func (n *TreeNode) insertBinarySearch(d int) {
	if n == nil {
		return
	}
	switch {
	case n.Val >= d:
		if n.Left == nil {
			n.Left = &TreeNode{Val: d}
		} else {
			n.Left.insertBinarySearch(d)
		}
	case n.Val < d:
		if n.Right == nil {
			n.Right = &TreeNode{Val: d}
		} else {
			n.Right.insertBinarySearch(d)
		}
	}
} */
