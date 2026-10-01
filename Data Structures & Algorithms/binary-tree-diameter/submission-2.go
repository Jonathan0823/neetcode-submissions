/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil { 
		return 0
	}

	leftHeight := maxHeight(root.Left)
	rightHeight := maxHeight(root.Right)
	diameter := leftHeight + rightHeight

	sub := max(diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right))

	return max(diameter, sub)
}

func maxHeight(cur *TreeNode) int { 
	if cur == nil { 
		return 0
	}

	return 1 + max(maxHeight(cur.Left), maxHeight(cur.Right))
}
