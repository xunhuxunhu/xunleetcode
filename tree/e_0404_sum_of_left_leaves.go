package tree

/*
Company: Google(2), Amazon(2), Expedia(2)

Find the sum of all left leaves in a given binary tree.

Example:
    3
   / \
  9  20
    /  \
   15   7
There are two left leaves in the binary tree, with values 9 and 15 respectively. Return 24.
*/

// top-down solution
func sumOfLeftLeaves404(root *TNode) int {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return 0
	}

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		return root.Left.Val + sumOfLeftLeaves404(root.Right)
	}

	return sumOfLeftLeaves404(root.Left) + sumOfLeftLeaves404(root.Right)
}