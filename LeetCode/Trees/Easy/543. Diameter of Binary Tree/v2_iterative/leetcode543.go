package main

import (
	"fmt"
	"os"

	"github.com/cloudemprise/leetcode543/tree"
)

var treeType string = "binary-tree"
var insertType string = "breadth-first"
var transversalType string = "Breadth-First"

/*
NOTE: The definition of the diameter of a binary tree as defined by this problem
is incorrect. The correct definition for the diameter of a binary tree is
defined as the number of NODES on the longest path between two end nodes and
not the number of EDGES as in the case of this problem.

Reference : https://www.geeksforgeeks.org/diameter-of-a-binary-tree/
*/

/*
Intuitive Understanding:
------------------------

The diameter of a tree T is the largest of the following quantities:

* the diameter of T’s left subtree.
* the diameter of T’s right subtree.
* the longest path between leaves that goes through the root of T, which can be
	computed from the heights of the subtrees of T.
*/

/*
The idea behind this solution is to use a modified depth-first search (DFS)
algorithm to traverse the binary tree in a post-order manner. We keep track of
the maximum diameter seen so far and the maximum depth of the left and right
subtrees of each node. We then update the maximum diameter seen so far with the
sum of the maximum depths of the left and right subtrees of each node.

Citations:
[1] https://stackoverflow.com/questions/74229104/iterative-postorder-traversal-of-binary-tree-in-python-optimality
[2] https://leetcode.com/problems/copy-list-with-random-pointer/discuss/43491/A-solution-with-constant-space-complexity-O(1)-and-linear-time-complexity-O(N)
[3] https://www.linkedin.com/pulse/grind-75-21-diameter-binary-tree-senthil-e
[4] https://youtube.com/watch?v=DcDjYOsjSlg
[5] https://www.geeksforgeeks.org/diameter-of-a-binary-tree/
[6] https://youtube.com/watch?v=qLcEd-1L1FE
*/

/*
The time complexity of this solution is O(n), where n is the number of nodes in
the binary tree. This is because we visit each node exactly once.

The space complexity of this solution is O(1), which is the space used by the
variables maxDiameter, stack, prev, leftMax, and rightMax. The maxDepth()
function is called only on the left and right children of each node, so its
space complexity is O(1) as well. Note: This solution is not as intuitive as the
other solutions and may be harder to understand and implement.
*/

func diameterOfBinaryTree(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}

	var maxDiameter int

	var stack []*tree.TreeNode
	var prev *tree.TreeNode
	var leftMax, rightMax int

	for root != nil || len(stack) > 0 {

		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		if root.Right == nil || root.Right == prev {
			stack = stack[:len(stack)-1]
			leftMax, rightMax = 0, 0

			if root.Left != nil {
				leftMax = maxDepth(root.Left) + 1
			}
			if root.Right != nil {
				rightMax = maxDepth(root.Right) + 1
			}

			maxDiameter = max(maxDiameter, leftMax+rightMax)
			prev = root
			root = nil
		} else {
			root = root.Right
		}
	}
	return maxDiameter
}

func maxDepth(root *tree.TreeNode) int {
	if root == nil {
		return -1
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	return 1 + max(leftDepth, rightDepth)
}

// max returns the larger of two integers.
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

///

func main() {

	var aTree = tree.New()

	const amount = 6
	var nums []int
	for i := range [amount]int{} {
		nums = append(nums, i+1)
	}

	fmt.Println("A 2D printout of a", treeType, "using a", insertType)
	fmt.Println("insertion algo is as follows: ")

	for _, num := range nums {
		fmt.Printf(">%v ", num)
		aTree.Insert(num)
	}
	fmt.Println()

	aTree.Print2dTree(os.Stdout)

	fmt.Println()
	fmt.Println("With a", transversalType, "transversal as follows:")
	fmt.Printf("[")
	aTree.PrintTree(os.Stdout)
	fmt.Println(" ]")
	fmt.Println()

	fmt.Println("Max Diameter = ", diameterOfBinaryTree(aTree.Root))

	fmt.Printf("\n//-------------------------------------\n\n")
}
