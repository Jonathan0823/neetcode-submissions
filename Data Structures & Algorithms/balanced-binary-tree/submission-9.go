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
	leftHeight := countHeight(root.Left)
	rightHeight := countHeight(root.Right)

	heightDif := leftHeight - rightHeight
	if heightDif < 0 { 
		heightDif *= -1
	}

	if heightDif >  1 { 
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func countHeight(node *TreeNode) int { 
	if node == nil { 
		return 0
	}

	return 1 + max(countHeight(node.Left), countHeight(node.Right))
}
