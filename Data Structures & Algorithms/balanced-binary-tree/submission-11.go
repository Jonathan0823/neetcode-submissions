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

func countHeight(node *TreeNode) int { 
	if node == nil { 
		return 0
	}

	leftHeight := countHeight(node.Left)
	if leftHeight == -1 { 
		return -1
	}

	rightHeight := countHeight(node.Right)
	if rightHeight == -1 { 
		return -1
	}

	difHeight := leftHeight - rightHeight
	if difHeight < 0 { 
		difHeight *= -1
	}

	if difHeight > 1 { 
		return -1
	}

	return 1 + max(leftHeight, rightHeight)
}
