/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return countHeight(root) != -1
}

func countHeight(root *TreeNode) int { 
	if root == nil { 
		return 0
	}

	leftHeight := countHeight(root.Left)
	if leftHeight == -1 {
		return -1 // Subtree kiri tidak seimbang, langsung return -1
	}
	rightHeight := countHeight(root.Right)
	if rightHeight == -1 {
		return -1 
	}

	difference := int(math.Abs(float64(leftHeight) - float64(rightHeight)))

	if difference > 1 { 
		return -1
	}

	return 1 + max(leftHeight, rightHeight)
}