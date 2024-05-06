package main

import (
	"fmt"
	"os"

	"github.com/cloudemprise/leetcode235/tree"
)

var treeType string = "binary-search-tree"
var insertType string = "manual"
var transversalType string = "Breadth-First"

/*
The function first checks if the root is nil. If it is, it returns nil.
Otherwise, it checks if both p and q are less than the root value. If they are,
it recursively calls lowestCommonAncestor on the left subtree of the root. If
both p and q are greater than the root value, it recursively calls
lowestCommonAncestor on the right subtree of the root. If neither of these
conditions are true, it means that p and q are on opposite sides of the root,
so the function returns the root.
*/

/*
The time complexity of this solution is O(h), where h is the height of the
binary search tree. In the worst case, the height of the tree is O(n), where n
is the number of nodes in the tree. Therefore, the worst-case time complexity is
O(n).

The space complexity of this solution is O(h), where h is the height of the
binary search tree. In the worst case, the height of the tree is O(n), where n
is the number of nodes in the tree. Therefore, the worst-case space complexity
is O(n).
*/

// lowestCommonAncestor recursively searches for the lowest common ancestor of
// two particular nodes of a given binary search tree.
func lowestCommonAncestor(root, p, q *tree.TreeNode) *tree.TreeNode {
	if root == nil {
		return nil
	}
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}

///

func main() {

	var aTree = tree.New()

	/*
					     6
					  /    \
					 2      8
					/ \    / \
				   0   4  7   8
		              / \
		             3   5
	*/

	aTree.Root = &tree.TreeNode{Val: 6,
		Left: &tree.TreeNode{Val: 2,
			Left: &tree.TreeNode{Val: 0},
			Right: &tree.TreeNode{Val: 4,
				Left:  &tree.TreeNode{Val: 3},
				Right: &tree.TreeNode{Val: 5},
			},
		},
		Right: &tree.TreeNode{Val: 8,
			Left:  &tree.TreeNode{Val: 7},
			Right: &tree.TreeNode{Val: 9},
		},
	}

	fmt.Println("A 2D printout of a", treeType, "using a", insertType)
	fmt.Println("insertion algo is as follows: ")

	fmt.Printf("\n")

	aTree.Print2dTree(os.Stdout)

	fmt.Println()
	fmt.Println("With a", transversalType, "transversal as follows:")
	fmt.Printf("[")
	aTree.PrintTree(os.Stdout)
	fmt.Println(" ]")

}
